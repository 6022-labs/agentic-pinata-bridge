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

		t.Run("Should be retryable on 429", func(t *testing.T) {
			t.Parallel()

			assert.True(t, apperrors.NewUpstreamStatusError("pinata", http.StatusTooManyRequests, "").IsRetryable())
		})

		t.Run("Should be retryable on 500", func(t *testing.T) {
			t.Parallel()

			assert.True(
				t,
				apperrors.NewUpstreamStatusError("pinata", http.StatusInternalServerError, "").IsRetryable(),
			)
		})

		t.Run("Should be retryable on 502", func(t *testing.T) {
			t.Parallel()

			assert.True(t, apperrors.NewUpstreamStatusError("pinata", http.StatusBadGateway, "").IsRetryable())
		})

		t.Run("Should be retryable on 503", func(t *testing.T) {
			t.Parallel()

			assert.True(
				t,
				apperrors.NewUpstreamStatusError("pinata", http.StatusServiceUnavailable, "").IsRetryable(),
			)
		})
	})

	t.Run("Given a refused request or a gateway timeout", func(t *testing.T) {
		t.Parallel()

		t.Run("Should not be retryable on 400", func(t *testing.T) {
			t.Parallel()

			assert.False(t, apperrors.NewUpstreamStatusError("pinata", http.StatusBadRequest, "").IsRetryable())
		})

		t.Run("Should not be retryable on 401", func(t *testing.T) {
			t.Parallel()

			assert.False(t, apperrors.NewUpstreamStatusError("pinata", http.StatusUnauthorized, "").IsRetryable())
		})

		t.Run("Should not be retryable on 404", func(t *testing.T) {
			t.Parallel()

			assert.False(t, apperrors.NewUpstreamStatusError("pinata", http.StatusNotFound, "").IsRetryable())
		})

		t.Run("Should not be retryable on 409", func(t *testing.T) {
			t.Parallel()

			assert.False(t, apperrors.NewUpstreamStatusError("pinata", http.StatusConflict, "").IsRetryable())
		})

		t.Run("Should not be retryable on 504", func(t *testing.T) {
			t.Parallel()

			assert.False(t, apperrors.NewUpstreamStatusError("pinata", http.StatusGatewayTimeout, "").IsRetryable())
		})
	})
}
