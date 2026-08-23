package metrics

import (
	"context"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

type ExternalHttpMetrics struct {
	durationHistogram     metric.Float64Histogram
	responseSizeHistogram metric.Int64Histogram
}

// NewExternalHttpMetrics builds the http.client.* instruments under the given meter scope.
// The convention defines no request/error counters: the duration histogram's count is the request
// count, and failures are the error.type dimension on it.
func NewExternalHttpMetrics(meterName string) *ExternalHttpMetrics {
	noopMeter := noop.NewMeterProvider().Meter(meterName)
	meter := otel.GetMeterProvider().Meter(meterName)

	durationHistogram, err := meter.Float64Histogram(
		"http.client.request.duration",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(LatencySecondsBuckets...),
		metric.WithDescription("Duration of HTTP client requests"),
	)
	if err != nil {
		durationHistogram, _ = noopMeter.Float64Histogram("http.client.request.duration")
	}

	responseSizeHistogram, err := meter.Int64Histogram(
		"http.client.response.body.size",
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(ResponseBodySizeBytesBuckets...),
		metric.WithDescription("Size of HTTP client response bodies"),
	)
	if err != nil {
		responseSizeHistogram, _ = noopMeter.Int64Histogram("http.client.response.body.size")
	}

	return &ExternalHttpMetrics{
		durationHistogram:     durationHistogram,
		responseSizeHistogram: responseSizeHistogram,
	}
}

// RecordRequest records one completed outbound request. serverUrl supplies semconv server.address
// and server.port; http.route is deliberately absent, being a server-side attribute.
func (m *ExternalHttpMetrics) RecordRequest(
	ctx context.Context,
	method, serverUrl string,
	statusCode int,
	duration time.Duration,
	responseSize int64,
) {
	attrs := []attribute.KeyValue{attribute.String("http.request.method", method)}
	attrs = append(attrs, ServerAttributes(serverUrl)...)
	attrs = append(attrs, attribute.Int("http.response.status_code", statusCode))
	// error.type is conditionally required, and only on a failed request.
	if statusCode >= 400 {
		attrs = append(attrs, attribute.String("error.type", strconv.Itoa(statusCode)))
	}

	option := metric.WithAttributes(attrs...)
	m.durationHistogram.Record(ctx, duration.Seconds(), option)
	if responseSize > 0 {
		m.responseSizeHistogram.Record(ctx, responseSize, option)
	}
}

// RecordTransportError records an outbound request that failed before any response.
func (m *ExternalHttpMetrics) RecordTransportError(
	ctx context.Context,
	method, serverUrl string,
	duration time.Duration,
	err error,
) {
	attrs := []attribute.KeyValue{attribute.String("http.request.method", method)}
	attrs = append(attrs, ServerAttributes(serverUrl)...)
	attrs = append(attrs, attribute.String("error.type", ErrorType(err)))

	m.durationHistogram.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
}
