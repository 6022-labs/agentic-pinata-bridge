package services

import (
	"context"
	"math/big"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/utils"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

// MintProposalImagesPinner pins proposal images up front so they resolve the moment the mint lands.
type MintProposalImagesPinner struct {
	logger                   *zap.Logger
	cidPinner                interfaces.CidPinnerInterface
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface
	pinMetrics               metrics_interfaces.PinMetricsInterface
}

func NewMintProposalImagesPinner(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *MintProposalImagesPinner {
	return &MintProposalImagesPinner{
		logger:                   logger,
		cidPinner:                cidPinner,
		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func (s *MintProposalImagesPinner) Pin(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	start := time.Now()
	defer func() {
		s.pinMetrics.RecordSweep(ctx, metrics_interfaces.SweepKindMintProposal, time.Since(start), err != nil)
	}()

	images, err := s.agentCollectionRequester.GetMintProposalImages(ctx, chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return errors.NewUnavailableError("mint_proposal_images_read_failed", upstreamFailureMessage)
	}

	var pinErrors error

	for _, image := range images {
		cid, ok := utils.ExtractCid(image)
		if !ok {
			s.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindMintProposal,
				metrics_interfaces.PinOutcomeInvalidCid,
			)
			s.logger.Warn("Mint proposal image is not a CID, skipping",
				zap.String("image", image),
				zap.Uint64("chainId", chainId),
				zap.String("agentCollectionAddress", agentCollectionAddress.String()),
				zap.String("proposalId", proposalId.String()),
			)
			continue
		}

		s.logger.Info("Pushing mint proposal image cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.String("proposalId", proposalId.String()),
		)

		// One unpinnable image must not cost the proposal its remaining ones.
		if err := s.cidPinner.Pin(ctx, cid); err != nil {
			s.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindMintProposal,
				metrics_interfaces.PinOutcomeFailed,
			)
			s.logger.Error("Failed to push mint proposal cid to pinata", zap.String("cid", cid), zap.Error(err))
			pinErrors = multierr.Append(pinErrors, err)
			continue
		}

		s.pinMetrics.RecordSweepImage(
			ctx,
			metrics_interfaces.SweepKindMintProposal,
			metrics_interfaces.PinOutcomePinned,
		)
	}

	if pinErrors != nil {
		return errors.NewUnavailableError("image_pin_failed", upstreamFailureMessage)
	}

	return nil
}
