package interfaces

import (
	"context"
	"time"
)

// ExternalHttpMetricsInterface records http.client.* metrics for an outbound
// HTTP adapter. Shared by every adapter that wraps an *http.Client transport.
type ExternalHttpMetricsInterface interface {
	RecordRequest(ctx context.Context, method, route string, statusCode int, duration time.Duration, responseSize int64)
	RecordTransportError(ctx context.Context, method, route string, duration time.Duration)
}
