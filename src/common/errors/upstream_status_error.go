package errors

import (
	"fmt"
	"net/http"
	"time"
)

// UpstreamStatusError is a dependency's unexpected HTTP status; it is internal, never an answer to a caller.
type UpstreamStatusError struct {
	Upstream   string
	StatusCode int
	RetryAfter time.Duration
	Body       string
}

func NewUpstreamStatusError(
	upstream string,
	statusCode int,
	retryAfter time.Duration,
	body string,
) *UpstreamStatusError {
	return &UpstreamStatusError{Upstream: upstream, StatusCode: statusCode, RetryAfter: retryAfter, Body: body}
}

func (e *UpstreamStatusError) Error() string {
	return fmt.Sprintf("%s returned %d: %s", e.Upstream, e.StatusCode, e.Body)
}

// IsRateLimited reports a 429: the dependency refused the call before working on it.
func (e *UpstreamStatusError) IsRateLimited() bool {
	return e.StatusCode == http.StatusTooManyRequests
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
