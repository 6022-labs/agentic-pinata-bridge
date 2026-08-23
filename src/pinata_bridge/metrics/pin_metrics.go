package metrics

import (
	"context"
	"time"

	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

const meterName = "agentic_pinata_bridge"

// errorTypeOther is the semconv catch-all for an error with no lower-cardinality classification.
const errorTypeOther = "_OTHER"

type PinMetrics struct {
	pinHistogram       metric.Float64Histogram
	hostLookupCounter  metric.Int64Counter
	hostLookupAttempts metric.Int64Histogram
	sweepHistogram     metric.Float64Histogram
	sweepImageCounter  metric.Int64Counter
}

func NewPinMetrics() *PinMetrics {
	noopMeter := noop.NewMeterProvider().Meter(meterName)
	meter := otel.GetMeterProvider().Meter(meterName)

	pinHistogram, err := meter.Float64Histogram(
		"pinata.pin.duration",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(common_metrics.LatencySecondsBuckets...),
		metric.WithDescription("Duration of pin attempts"),
	)
	if err != nil {
		pinHistogram, _ = noopMeter.Float64Histogram("pinata.pin.duration")
	}

	hostLookupCounter, err := meter.Int64Counter(
		"ipfs.host_lookup.count",
		metric.WithDescription("Total number of ipfs-check host lookups by outcome"),
	)
	if err != nil {
		hostLookupCounter, _ = noopMeter.Int64Counter("ipfs.host_lookup.count")
	}

	hostLookupAttempts, err := meter.Int64Histogram(
		"ipfs.host_lookup.attempts",
		metric.WithUnit("{attempt}"),
		metric.WithExplicitBucketBoundaries(common_metrics.RetryAttemptsBuckets...),
		metric.WithDescription("Attempts spent per ipfs-check host lookup"),
	)
	if err != nil {
		hostLookupAttempts, _ = noopMeter.Int64Histogram("ipfs.host_lookup.attempts")
	}

	sweepHistogram, err := meter.Float64Histogram(
		"pinata.sweep.duration",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(common_metrics.LongLatencySecondsBuckets...),
		metric.WithDescription("Duration of a backfill sweep"),
	)
	if err != nil {
		sweepHistogram, _ = noopMeter.Float64Histogram("pinata.sweep.duration")
	}

	sweepImageCounter, err := meter.Int64Counter(
		"pinata.sweep.images",
		metric.WithDescription("Images visited by a sweep, by outcome"),
	)
	if err != nil {
		sweepImageCounter, _ = noopMeter.Int64Counter("pinata.sweep.images")
	}

	return &PinMetrics{
		pinHistogram:       pinHistogram,
		hostLookupCounter:  hostLookupCounter,
		hostLookupAttempts: hostLookupAttempts,
		sweepHistogram:     sweepHistogram,
		sweepImageCounter:  sweepImageCounter,
	}
}

func (m *PinMetrics) RecordPin(ctx context.Context, outcome string, withHostAddresses bool, duration time.Duration) {
	attrs := metric.WithAttributes(
		attribute.String("outcome", outcome),
		attribute.Bool("pinata.host_addresses", withHostAddresses),
	)
	m.pinHistogram.Record(ctx, duration.Seconds(), attrs)
}

func (m *PinMetrics) RecordHostLookup(ctx context.Context, outcome string, attempts int64) {
	attrs := metric.WithAttributes(attribute.String("outcome", outcome))
	m.hostLookupCounter.Add(ctx, 1, attrs)
	m.hostLookupAttempts.Record(ctx, attempts, attrs)
}

func (m *PinMetrics) RecordSweep(ctx context.Context, kind string, duration time.Duration, failed bool) {
	attrs := []attribute.KeyValue{attribute.String("pinata.sweep.kind", kind)}
	// error.type is conditionally required, and only on a failed operation.
	if failed {
		attrs = append(attrs, attribute.String("error.type", errorTypeOther))
	}
	m.sweepHistogram.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
}

func (m *PinMetrics) RecordSweepImage(ctx context.Context, kind, outcome string) {
	m.sweepImageCounter.Add(ctx, 1, metric.WithAttributes(
		attribute.String("pinata.sweep.kind", kind),
		attribute.String("outcome", outcome),
	))
}
