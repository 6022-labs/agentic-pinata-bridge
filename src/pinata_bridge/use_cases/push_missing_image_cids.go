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

// upstreamFailureMessage keeps upstream text (RPC urls carry api keys) off the wire; the log has the detail.
const upstreamFailureMessage = "upstream request failed"

// PushMissingImageCids sweeps every configured chain and pins whatever Pinata is missing.
type PushMissingImageCids struct {
	logger                           *zap.Logger
	agentCollectionRequester         interfaces.AgentCollectionRequesterInterface
	pinMetrics                       metrics_interfaces.PinMetricsInterface
	chainsSettings                   *settings.ChainsSettings
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface
	agentImagesPinner                interfaces.AgentImagesPinnerInterface
	pinTracer                        traces_interfaces.PinTracerInterface
}

func NewPushMissingImageCids(
	logger *zap.Logger,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
	chainsSettings *settings.ChainsSettings,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	agentImagesPinner interfaces.AgentImagesPinnerInterface,
	pinTracer traces_interfaces.PinTracerInterface,
) *PushMissingImageCids {
	return &PushMissingImageCids{
		logger:                           logger,
		agentCollectionRequester:         agentCollectionRequester,
		pinMetrics:                       pinMetrics,
		chainsSettings:                   chainsSettings,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentImagesPinner:                agentImagesPinner,
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

	start := time.Now()
	defer func() { u.pinMetrics.RecordSweep(ctx, metrics_interfaces.SweepKindAll, time.Since(start), err != nil) }()

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
				if err := u.agentImagesPinner.PinMissing(ctx, chainId, collectionAddress, tokenId); err != nil {
					u.logger.Error("Failed to push agent image cid to pinata", zap.Error(err))
					continue
				}

				u.logger.Info("Successfully pushed agent image cid to pinata", zap.String("tokenId", tokenId.String()))
			}
		}
	}

	return &responses.PushResponse{}, nil
}
