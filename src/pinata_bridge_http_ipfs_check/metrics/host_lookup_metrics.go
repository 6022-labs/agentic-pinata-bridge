package metrics

import (
	"context"

	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

const meterName = "agentic_pinata_bridge_http_ipfs_check"

type HostLookupMetrics struct {
	lookupCounter  metric.Int64Counter
	lookupAttempts metric.Int64Histogram
}

func NewHostLookupMetrics() *HostLookupMetrics {
	noopMeter := noop.NewMeterProvider().Meter(meterName)
	meter := otel.GetMeterProvider().Meter(meterName)

	lookupCounter, err := meter.Int64Counter(
		"ipfs.host_lookup.count",
		metric.WithDescription("Total number of ipfs-check host lookups by outcome"),
	)
	if err != nil {
		lookupCounter, _ = noopMeter.Int64Counter("ipfs.host_lookup.count")
	}

	lookupAttempts, err := meter.Int64Histogram(
		"ipfs.host_lookup.attempts",
		metric.WithUnit("{attempt}"),
		metric.WithExplicitBucketBoundaries(common_metrics.RetryAttemptsBuckets...),
		metric.WithDescription("Attempts spent per ipfs-check host lookup"),
	)
	if err != nil {
		lookupAttempts, _ = noopMeter.Int64Histogram("ipfs.host_lookup.attempts")
	}

	return &HostLookupMetrics{lookupCounter: lookupCounter, lookupAttempts: lookupAttempts}
}

func (m *HostLookupMetrics) RecordHostLookup(ctx context.Context, outcome string, attempts int64) {
	attrs := metric.WithAttributes(attribute.String("outcome", outcome))
	m.lookupCounter.Add(ctx, 1, attrs)
	m.lookupAttempts.Record(ctx, attempts, attrs)
}
