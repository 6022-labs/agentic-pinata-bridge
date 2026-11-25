package errors

import (
	"fmt"
	"net/http"
)

// UpstreamStatusError is a dependency's unexpected HTTP status; it is internal, never an answer to a caller.
type UpstreamStatusError struct {
	Upstream   string
	StatusCode int
	Body       string
}

func NewUpstreamStatusError(upstream string, statusCode int, body string) *UpstreamStatusError {
	return &UpstreamStatusError{Upstream: upstream, StatusCode: statusCode, Body: body}
}

func (e *UpstreamStatusError) Error() string {
	return fmt.Sprintf("%s returned %d: %s", e.Upstream, e.StatusCode, e.Body)
}

// IsRetryable leaves 504 and other gateway timeouts out, like any timeout.
func (e *UpstreamStatusError) IsRetryable() bool {
	switch e.StatusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}
