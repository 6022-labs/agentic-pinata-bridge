package use_cases

import (
	"context"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"go.uber.org/zap"
)

// AbstractPushUseCase carries the collaborators every push use case shares.
type AbstractPushUseCase struct {
	logger                   *zap.Logger
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface
	pinMetrics               metrics_interfaces.PinMetricsInterface
}

func NewAbstractPushUseCase(
	logger *zap.Logger,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *AbstractPushUseCase {
	return &AbstractPushUseCase{
		logger:                   logger,
		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func (u *AbstractPushUseCase) recordSweep(ctx context.Context, kind string, start time.Time, err *error) {
	u.pinMetrics.RecordSweep(ctx, kind, time.Since(start), err != nil && *err != nil)
}
