package use_cases

import (
	"context"
	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"math/big"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// PushImagesOfMintProposal pins every image carried by one mint proposal.
type PushImagesOfMintProposal struct {
	*AbstractPushUseCase

	cidPinner interfaces.CidPinnerInterface
}

func NewPushImagesOfMintProposal(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *PushImagesOfMintProposal {
	return &PushImagesOfMintProposal{
		AbstractPushUseCase: NewAbstractPushUseCase(logger, agentCollectionRequester, pinMetrics),
		cidPinner:           cidPinner,
	}
}

func (u *PushImagesOfMintProposal) Execute(
	ctx context.Context,
	request *requests.PushImagesOfMintProposalRequest,
) (*responses.PushResponse, error) {
	if err := request.ValidateAndSanitize(); err != nil {
		return nil, err
	}

	if err := u.push(
		ctx,
		request.ChainIdValue(),
		request.AgentCollectionAddressValue(),
		request.ProposalIdValue(),
	); err != nil {
		return nil, err
	}

	return &responses.PushResponse{}, nil
}

// push is the chain-typed entry point the mint-proposal-created use case reuses.
func (u *PushImagesOfMintProposal) push(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	defer u.recordSweep(ctx, metrics_interfaces.SweepKindMintProposal, time.Now(), &err)

	cids, err := u.agentCollectionRequester.GetMintProposalImages(ctx, chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return errors.NewUnavailableError("mint_proposal_images_read_failed", upstreamFailureMessage)
	}

	for _, cid := range cids {
		u.logger.Info("Pushing mint proposal image cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("proposalId", proposalId.Int64()),
		)

		if err := u.cidPinner.Pin(ctx, cid); err != nil {
			u.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindMintProposal,
				metrics_interfaces.PinOutcomeFailed,
			)

			u.logger.Error("Failed to push mint proposal cid to pinata", zap.String("cid", cid), zap.Error(err))

			return errors.NewUnavailableError("image_pin_failed", upstreamFailureMessage)
		}

		u.pinMetrics.RecordSweepImage(
			ctx,
			metrics_interfaces.SweepKindMintProposal,
			metrics_interfaces.PinOutcomePinned,
		)
	}

	return nil
}
