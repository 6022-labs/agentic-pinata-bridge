package use_cases

import (
	"context"
	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	traces_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
	"go.uber.org/zap"
)

// PushMissingImageCids sweeps every configured chain and pins whatever Pinata is missing.
type PushMissingImageCids struct {
	*AbstractPushUseCase

	chainsSettings                   *settings.ChainsSettings
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface
	pushMissingImagesOfAgent         *PushMissingImagesOfAgent
	pinTracer                        traces_interfaces.PinTracerInterface
}

func NewPushMissingImageCids(
	logger *zap.Logger,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
	chainsSettings *settings.ChainsSettings,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	pushMissingImagesOfAgent *PushMissingImagesOfAgent,
	pinTracer traces_interfaces.PinTracerInterface,
) *PushMissingImageCids {
	return &PushMissingImageCids{
		AbstractPushUseCase:              NewAbstractPushUseCase(logger, agentCollectionRequester, pinMetrics),
		chainsSettings:                   chainsSettings,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pushMissingImagesOfAgent:         pushMissingImagesOfAgent,
		pinTracer:                        pinTracer,
	}
}

func (u *PushMissingImageCids) Execute(ctx context.Context) (response *responses.PushResponse, err error) {
	ctx, span := u.pinTracer.StartSweep(ctx, metrics_interfaces.SweepKindAll)
	defer span.End()
	defer func() {
		if err != nil {
			span.Fail(err)
		}
	}()

	defer u.recordSweep(ctx, metrics_interfaces.SweepKindAll, time.Now(), &err)

	for _, chainId := range u.chainsSettings.ChainIds() {
		u.logger.Info("Processing chain", zap.Uint64("chainId", chainId))

		allCollections, err := u.agentCollectionsManagerRequester.GetAllCollectionAddresses(ctx, chainId)
		if err != nil {
			return nil, errors.NewUnavailableError("collections_read_failed", upstreamFailureMessage)
		}

		for _, collectionAddress := range allCollections {
			u.logger.Info("Processing collection",
				zap.Uint64("chainId", chainId),
				zap.String("collectionAddress", collectionAddress.String()),
			)

			tokenIds, err := u.agentCollectionRequester.GetAllTokenIds(ctx, chainId, collectionAddress)
			if err != nil {
				return nil, errors.NewUnavailableError("token_ids_read_failed", upstreamFailureMessage)
			}

			for _, tokenId := range tokenIds {
				if err := u.pushMissingImagesOfAgent.push(ctx, chainId, collectionAddress, tokenId); err != nil {
					u.logger.Error("Failed to push agent image cid to pinata", zap.Error(err))
					continue
				}

				u.logger.Info("Successfully pushed agent image cid to pinata", zap.String("tokenId", tokenId.String()))
			}
		}
	}

	return &responses.PushResponse{}, nil
}
