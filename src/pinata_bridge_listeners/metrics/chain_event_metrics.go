package metrics

import (
	"context"
	"strconv"
	"time"

	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

const meterName = "agentic_pinata_bridge_listeners"

type ChainEventMetrics struct {
	handleHistogram     metric.Float64Histogram
	activeSubscriptions metric.Int64UpDownCounter
	subscriptionErrors  metric.Int64Counter
}

func NewChainEventMetrics() *ChainEventMetrics {
	noopMeter := noop.NewMeterProvider().Meter(meterName)
	meter := otel.GetMeterProvider().Meter(meterName)

	handleHistogram, err := meter.Float64Histogram(
		"chain.event.handle.duration",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(common_metrics.LongLatencySecondsBuckets...),
		metric.WithDescription("Duration of chain event handling"),
	)
	if err != nil {
		handleHistogram, _ = noopMeter.Float64Histogram("chain.event.handle.duration")
	}

	activeSubscriptions, err := meter.Int64UpDownCounter(
		"chain.subscription.active",
		metric.WithDescription("Currently open chain event subscriptions"),
	)
	if err != nil {
		activeSubscriptions, _ = noopMeter.Int64UpDownCounter("chain.subscription.active")
	}

	subscriptionErrors, err := meter.Int64Counter(
		"chain.subscription.errors",
		metric.WithDescription("Total number of chain subscription errors"),
	)
	if err != nil {
		subscriptionErrors, _ = noopMeter.Int64Counter("chain.subscription.errors")
	}

	return &ChainEventMetrics{
		handleHistogram:     handleHistogram,
		activeSubscriptions: activeSubscriptions,
		subscriptionErrors:  subscriptionErrors,
	}
}

func (m *ChainEventMetrics) RecordEvent(
	ctx context.Context,
	eventName string,
	chainId uint64,
	outcome string,
	duration time.Duration,
) {
	attrs := metric.WithAttributes(
		attribute.String("chain.event.name", eventName),
		attribute.String("chain.id", strconv.FormatUint(chainId, 10)),
		attribute.String("outcome", outcome),
	)
	m.handleHistogram.Record(ctx, duration.Seconds(), attrs)
}

func (m *ChainEventMetrics) RecordSubscriptionOpened(ctx context.Context, eventName string, chainId uint64) {
	m.activeSubscriptions.Add(ctx, 1, m.subscriptionAttributes(eventName, chainId))
}

func (m *ChainEventMetrics) RecordSubscriptionClosed(ctx context.Context, eventName string, chainId uint64) {
	m.activeSubscriptions.Add(ctx, -1, m.subscriptionAttributes(eventName, chainId))
}

func (m *ChainEventMetrics) RecordSubscriptionError(ctx context.Context, eventName string, chainId uint64) {
	m.subscriptionErrors.Add(ctx, 1, m.subscriptionAttributes(eventName, chainId))
}

func (m *ChainEventMetrics) subscriptionAttributes(eventName string, chainId uint64) metric.MeasurementOption {
	return metric.WithAttributes(
		attribute.String("chain.event.name", eventName),
		attribute.String("chain.id", strconv.FormatUint(chainId, 10)),
	)
}
