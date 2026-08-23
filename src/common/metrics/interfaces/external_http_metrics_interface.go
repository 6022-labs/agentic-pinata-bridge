package interfaces

import (
	"context"
	"time"
)

// ExternalHttpMetricsInterface records http.client.* metrics for outbound calls.
type ExternalHttpMetricsInterface interface {
	RecordRequest(ctx context.Context, method, serverUrl string, statusCode int, duration time.Duration, responseSize int64)
	RecordTransportError(ctx context.Context, method, serverUrl string, duration time.Duration, err error)
}
