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

// upstreamFailureMessage keeps upstream text (RPC urls carry api keys) off the wire; the log has the detail.
const upstreamFailureMessage = "upstream request failed"

// AgentImagesPinner is the one place an agent's images get swept, whatever triggered it.
type AgentImagesPinner struct {
	logger                   *zap.Logger
	cidPinner                interfaces.CidPinnerInterface
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface
	pinataRequester          interfaces.PinataRequesterInterface
	pinMetrics               metrics_interfaces.PinMetricsInterface
}

func NewAgentImagesPinner(
	logger *zap.Logger,
	cidPinner interfaces.CidPinnerInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinataRequester interfaces.PinataRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *AgentImagesPinner {
	return &AgentImagesPinner{
		logger:                   logger,
		cidPinner:                cidPinner,
		agentCollectionRequester: agentCollectionRequester,
		pinataRequester:          pinataRequester,
		pinMetrics:               pinMetrics,
	}
}

func (s *AgentImagesPinner) PinMissing(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	tokenId big.Int,
) (err error) {
	start := time.Now()
	defer func() {
		s.pinMetrics.RecordSweep(ctx, metrics_interfaces.SweepKindAgent, time.Since(start), err != nil)
	}()

	images, err := s.agentCollectionRequester.GetAgentImages(ctx, chainId, agentCollectionAddress, tokenId)
	if err != nil {
		s.logger.Error("Failed to read agent images", zap.Uint64("chainId", chainId), zap.Error(err))

		return errors.NewUnavailableError("agent_images_read_failed", upstreamFailureMessage)
	}

	var pinErrors error

	for _, image := range images {
		cid, ok := utils.ExtractCid(image)
		if !ok {
			s.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindAgent,
				metrics_interfaces.PinOutcomeInvalidCid,
			)
			s.logger.Warn("Agent image is not a CID, skipping",
				zap.String("image", image),
				zap.Uint64("chainId", chainId),
				zap.String("agentCollectionAddress", agentCollectionAddress.String()),
				zap.String("agentCollectionTokenId", tokenId.String()),
			)
			continue
		}

		isUploaded, err := s.pinataRequester.IsCidUploaded(ctx, cid)
		if err != nil {
			s.logger.Error("Failed to check if cid is uploaded", zap.String("cid", cid), zap.Error(err))

			return errors.NewUnavailableError("pin_status_read_failed", upstreamFailureMessage)
		}

		if *isUploaded {
			s.pinMetrics.RecordSweepImage(
				ctx,
				metrics_interfaces.SweepKindAgent,
				metrics_interfaces.PinOutcomeAlreadyPinned,
			)
			s.logger.Debug("CID already uploaded to pinata, skipping", zap.String("cid", cid))
			continue
		}

		s.logger.Info("Pushing agent images cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.String("agentCollectionTokenId", tokenId.String()),
		)

		// One unpinnable image must not cost the agent its remaining ones.
		if err := s.cidPinner.Pin(ctx, cid); err != nil {
			s.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			s.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			pinErrors = multierr.Append(pinErrors, err)
			continue
		}

		s.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
	}

	if pinErrors != nil {
		return errors.NewUnavailableError("image_pin_failed", upstreamFailureMessage)
	}

	return nil
}
