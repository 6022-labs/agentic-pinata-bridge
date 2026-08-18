package use_cases

import (
	"context"
	"math/big"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// PushImageOfAgentImageProposal pins the single image carried by one agent-image proposal.
type PushImageOfAgentImageProposal struct {
	*AbstractPushUseCase

	cidPinner interfaces.CidPinnerInterface
}

func NewPushImageOfAgentImageProposal(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *PushImageOfAgentImageProposal {
	return &PushImageOfAgentImageProposal{
		AbstractPushUseCase: NewAbstractPushUseCase(logger, agentCollectionRequester, pinMetrics),
		cidPinner:           cidPinner,
	}
}

func (u *PushImageOfAgentImageProposal) Execute(
	ctx context.Context,
	request *requests.PushImageOfAgentImageProposalRequest,
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

// push is the chain-typed entry point the image-proposal-created use case reuses.
func (u *PushImageOfAgentImageProposal) push(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	defer u.recordSweep(ctx, metrics_interfaces.SweepKindImageProposal, time.Now(), &err)

	cid, err := u.agentCollectionRequester.GetAgentImageProposalImage(ctx, chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return err
	}

	u.logger.Info("Pushing agent image proposal cid to pinata",
		zap.String("cid", *cid),
		zap.String("agentCollectionAddress", agentCollectionAddress.String()),
		zap.Int64("proposalId", proposalId.Int64()),
	)

	if err := u.cidPinner.Pin(ctx, *cid); err != nil {
		u.pinMetrics.RecordSweepImage(
			ctx,
			metrics_interfaces.SweepKindImageProposal,
			metrics_interfaces.PinOutcomeFailed,
		)
		u.logger.Error("Failed to push agent image proposal cid to pinata", zap.String("cid", *cid), zap.Error(err))
		return err
	}

	u.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomePinned)

	return nil
}
