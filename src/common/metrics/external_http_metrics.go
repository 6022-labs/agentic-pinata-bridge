package metrics

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

type ExternalHttpMetrics struct {
	requestCounter        metric.Int64Counter
	durationHistogram     metric.Float64Histogram
	responseSizeHistogram metric.Int64Histogram
	errorCounter          metric.Int64Counter
}

// NewExternalHttpMetrics builds the http.client.* instruments under the given meter scope.
func NewExternalHttpMetrics(meterName string) *ExternalHttpMetrics {
	noopMeter := noop.NewMeterProvider().Meter(meterName)
	meter := otel.GetMeterProvider().Meter(meterName)

	requestCounter, err := meter.Int64Counter(
		"http.client.request.count",
		metric.WithDescription("Total number of outbound HTTP requests"),
	)
	if err != nil {
		requestCounter, _ = noopMeter.Int64Counter("http.client.request.count")
	}

	durationHistogram, err := meter.Float64Histogram(
		"http.client.request.duration",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(LatencySecondsBuckets...),
		metric.WithDescription("Duration of outbound HTTP requests"),
	)
	if err != nil {
		durationHistogram, _ = noopMeter.Float64Histogram("http.client.request.duration")
	}

	responseSizeHistogram, err := meter.Int64Histogram(
		"http.client.response.body.size",
		metric.WithUnit("By"),
		metric.WithDescription("Response size of outbound HTTP requests"),
	)
	if err != nil {
		responseSizeHistogram, _ = noopMeter.Int64Histogram("http.client.response.body.size")
	}

	errorCounter, err := meter.Int64Counter(
		"http.client.request.errors",
		metric.WithDescription("Total number of outbound HTTP request errors"),
	)
	if err != nil {
		errorCounter, _ = noopMeter.Int64Counter("http.client.request.errors")
	}

	return &ExternalHttpMetrics{
		requestCounter:        requestCounter,
		durationHistogram:     durationHistogram,
		responseSizeHistogram: responseSizeHistogram,
		errorCounter:          errorCounter,
	}
}

// RecordRequest records one completed outbound request: count, duration, response
// size (when >0), and a 5xx error. responseSize <= 0 is omitted.
func (m *ExternalHttpMetrics) RecordRequest(
	ctx context.Context,
	method, route string,
	statusCode int,
	duration time.Duration,
	responseSize int64,
) {
	attrs := metric.WithAttributes(
		attribute.String("http.request.method", method),
		attribute.String("http.route", route),
		attribute.Int("http.response.status_code", statusCode),
	)
	m.requestCounter.Add(ctx, 1, attrs)
	m.durationHistogram.Record(ctx, duration.Seconds(), attrs)
	if responseSize > 0 {
		m.responseSizeHistogram.Record(ctx, responseSize, attrs)
	}
	if statusCode >= 500 {
		m.errorCounter.Add(ctx, 1, attrs)
	}
}

// RecordTransportError records an outbound request that failed before a response
// (DNS, timeout, connection refused): count, duration, and an error, with no status.
func (m *ExternalHttpMetrics) RecordTransportError(ctx context.Context, method, route string, duration time.Duration) {
	attrs := metric.WithAttributes(
		attribute.String("http.request.method", method),
		attribute.String("http.route", route),
	)
	m.requestCounter.Add(ctx, 1, attrs)
	m.durationHistogram.Record(ctx, duration.Seconds(), attrs)
	m.errorCounter.Add(ctx, 1, attrs)
}
