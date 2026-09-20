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
	"go.uber.org/zap"
)

// AgentImageProposalImagePinner pins a proposed replacement image before a moderator applies it.
type AgentImageProposalImagePinner struct {
	logger                   *zap.Logger
	cidPinner                interfaces.CidPinnerInterface
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface
	pinMetrics               metrics_interfaces.PinMetricsInterface
}

func NewAgentImageProposalImagePinner(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *AgentImageProposalImagePinner {
	return &AgentImageProposalImagePinner{
		logger:                   logger,
		cidPinner:                cidPinner,
		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func (s *AgentImageProposalImagePinner) Pin(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	start := time.Now()
	defer func() {
		s.pinMetrics.RecordSweep(ctx, metrics_interfaces.SweepKindImageProposal, time.Since(start), err != nil)
	}()

	image, err := s.agentCollectionRequester.GetAgentImageProposalImage(
		ctx,
		chainId,
		agentCollectionAddress,
		proposalId,
	)
	if err != nil {
		return errors.NewUnavailableError("image_proposal_read_failed", upstreamFailureMessage)
	}

	cid, ok := utils.ExtractCid(*image)
	if !ok {
		s.pinMetrics.RecordSweepImage(
			ctx,
			metrics_interfaces.SweepKindImageProposal,
			metrics_interfaces.PinOutcomeInvalidCid,
		)
		s.logger.Warn("Agent image proposal image is not a CID, skipping",
			zap.String("image", *image),
			zap.Uint64("chainId", chainId),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.String("proposalId", proposalId.String()),
		)

		return nil
	}

	s.logger.Info("Pushing agent image proposal cid to pinata",
		zap.String("cid", cid),
		zap.String("agentCollectionAddress", agentCollectionAddress.String()),
		zap.String("proposalId", proposalId.String()),
	)

	if err := s.cidPinner.Pin(ctx, cid); err != nil {
		s.pinMetrics.RecordSweepImage(
			ctx,
			metrics_interfaces.SweepKindImageProposal,
			metrics_interfaces.PinOutcomeFailed,
		)
		s.logger.Error("Failed to push agent image proposal cid to pinata", zap.String("cid", cid), zap.Error(err))

		return errors.NewUnavailableError("image_pin_failed", upstreamFailureMessage)
	}

	s.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomePinned)

	return nil
}
