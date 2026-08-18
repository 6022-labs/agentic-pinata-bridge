package metrics_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"github.com/stretchr/testify/assert"
)

func TestWhenTemplatizingARoute(t *testing.T) {
	t.Parallel()

	t.Run("Given a path with a UUID segment", func(t *testing.T) {
		t.Parallel()

		t.Run("Should collapse the UUID to {id}", func(t *testing.T) {
			t.Parallel()

			out := metrics.TemplatizeRoute("/internal/config/mcp/4f9a1c2e-2b6d-4c1a-9f3e-0a1b2c3d4e5f")

			assert.Equal(t, "/internal/config/mcp/{id}", out)
		})
	})

	t.Run("Given a path with a numeric segment", func(t *testing.T) {
		t.Parallel()

		t.Run("Should collapse the number to {id}", func(t *testing.T) {
			t.Parallel()

			out := metrics.TemplatizeRoute("/internal/conversations/12345/messages")

			assert.Equal(t, "/internal/conversations/{id}/messages", out)
		})
	})

	t.Run("Given a path with a long hex segment", func(t *testing.T) {
		t.Parallel()

		t.Run("Should collapse the hex blob to {id}", func(t *testing.T) {
			t.Parallel()

			out := metrics.TemplatizeRoute("/tx/0xabcdef0123456789abcdef")

			assert.Equal(t, "/tx/{id}", out)
		})
	})

	t.Run("Given a fully static path", func(t *testing.T) {
		t.Parallel()

		t.Run("Should leave it unchanged", func(t *testing.T) {
			t.Parallel()

			out := metrics.TemplatizeRoute("/internal/responses")

			assert.Equal(t, "/internal/responses", out)
		})
	})

	t.Run("Given an empty path", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return it unchanged", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "", metrics.TemplatizeRoute(""))
		})
	})
}
