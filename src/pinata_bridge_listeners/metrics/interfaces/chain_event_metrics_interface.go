package interfaces

import (
	"context"
	"time"
)

const (
	ChainEventOutcomeHandled = "handled"
	ChainEventOutcomeFailed  = "failed"
)

// ChainEventMetricsInterface records the chain.* metrics shared by every 6022 chain-event listener.
type ChainEventMetricsInterface interface {
	RecordEvent(ctx context.Context, eventName string, chainId uint64, outcome string, duration time.Duration)
	RecordSubscriptionOpened(ctx context.Context, eventName string, chainId uint64)
	RecordSubscriptionClosed(ctx context.Context, eventName string, chainId uint64)
	RecordSubscriptionError(ctx context.Context, eventName string, chainId uint64)
}
