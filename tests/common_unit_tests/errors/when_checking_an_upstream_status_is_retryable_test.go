package errors_test

import (
	"net/http"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/stretchr/testify/assert"
)

func TestWhenCheckingAnUpstreamStatusIsRetryable(t *testing.T) {
	t.Parallel()

	t.Run("Given a rate limit or a server error", func(t *testing.T) {
		t.Parallel()

		for _, status := range []int{http.StatusTooManyRequests, http.StatusInternalServerError, http.StatusBadGateway} {
			t.Run("Should be retryable on "+http.StatusText(status), func(t *testing.T) {
				t.Parallel()

				assert.True(t, apperrors.NewUpstreamStatusError("pinata", status, "").IsRetryable())
			})
		}
	})

	t.Run("Given a refused request", func(t *testing.T) {
		t.Parallel()

		for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound} {
			t.Run("Should not be retryable on "+http.StatusText(status), func(t *testing.T) {
				t.Parallel()

				assert.False(t, apperrors.NewUpstreamStatusError("pinata", status, "").IsRetryable())
			})
		}
	})
}
