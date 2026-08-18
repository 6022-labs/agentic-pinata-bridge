package use_cases

import (
	"context"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/big"
)

// PushMissingImagesOfAgent pins every image of one agent that Pinata does not already hold.
type PushMissingImagesOfAgent struct {
	*AbstractPushUseCase

	pinataRequester interfaces.PinataRequesterInterface
}

func NewPushMissingImagesOfAgent(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinataRequester interfaces.PinataRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *PushMissingImagesOfAgent {
	return &PushMissingImagesOfAgent{
		AbstractPushUseCase: NewAbstractPushUseCase(logger, cidPinner, agentCollectionRequester, pinMetrics),
		pinataRequester:     pinataRequester,
	}
}

func (u *PushMissingImagesOfAgent) Execute(
	ctx context.Context,
	request *requests.PushMissingImagesOfAgentRequest,
) (response *responses.PushResponse, err error) {
	if err := request.ValidateAndSanitize(); err != nil {
		return nil, err
	}

	if err := u.push(
		ctx,
		request.ChainIdValue(),
		request.AgentCollectionAddressValue(),
		request.AgentCollectionTokenIdValue(),
	); err != nil {
		return nil, err
	}

	return &responses.PushResponse{}, nil
}

// push is the chain-typed entry point the minted-event use case reuses.
func (u *PushMissingImagesOfAgent) push(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	agentCollectionTokenId big.Int,
) (err error) {
	defer u.recordSweep(ctx, metrics_interfaces.SweepKindAgent, time.Now(), &err)

	cids, err := u.agentCollectionRequester.GetAgentImages(ctx, chainId, agentCollectionAddress, agentCollectionTokenId)
	if err != nil {
		return err
	}

	for _, cid := range cids {
		isUploaded, err := u.pinataRequester.IsCidUploaded(ctx, cid)
		if err != nil {
			u.logger.Error("Failed to check if cid is uploaded", zap.String("cid", cid), zap.Error(err))
			return err
		}

		if *isUploaded {
			u.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindAgent,
				metrics_interfaces.PinOutcomeAlreadyPinned,
			)
			u.logger.Debug("CID already uploaded to pinata, skipping", zap.String("cid", cid))
			continue
		}

		u.logger.Info("Pushing agent images cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("agentCollectionTokenId", agentCollectionTokenId.Int64()),
		)

		if err := u.cidPinner.Pin(ctx, cid); err != nil {
			u.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			u.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			return err
		}

		u.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
	}

	return nil
}
