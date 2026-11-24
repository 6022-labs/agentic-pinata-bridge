package settings_test

import (
	"testing"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/settings"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func panicLogger() *zap.Logger {
	return zap.NewNop().WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic))
}

func TestWhenLoadingIpfsCheckSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given a configured base url", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url": "  https://ipfs-check.example.com  ",
		}, "."), nil)

		t.Run("Should load and trim the base url", func(t *testing.T) {
			t.Parallel()

			result := settings.NewIpfsCheckSettings(zap.NewNop(), k)

			assert.Equal(t, "https://ipfs-check.example.com", result.BaseUrl)
		})
	})

	t.Run("Given the base url is missing", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url": "   ",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewIpfsCheckSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given the request timeout is configured", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url":                "https://ipfs-check.example.com",
			"ipfs_check.request_timeout_seconds": 5,
		}, "."), nil)

		t.Run("Should use it", func(t *testing.T) {
			t.Parallel()

			result := settings.NewIpfsCheckSettings(zap.NewNop(), k)

			assert.Equal(t, 5*time.Second, result.RequestTimeout)
		})
	})

	t.Run("Given the request timeout is not configured", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url": "https://ipfs-check.example.com",
		}, "."), nil)

		t.Run("Should default to 90 seconds", func(t *testing.T) {
			t.Parallel()

			result := settings.NewIpfsCheckSettings(zap.NewNop(), k)

			assert.Equal(t, 90*time.Second, result.RequestTimeout)
		})
	})

	t.Run("Given the request timeout is zero", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url":                "https://ipfs-check.example.com",
			"ipfs_check.request_timeout_seconds": 0,
		}, "."), nil)

		t.Run("Should fall back to 90 seconds", func(t *testing.T) {
			t.Parallel()

			result := settings.NewIpfsCheckSettings(zap.NewNop(), k)

			assert.Equal(t, 90*time.Second, result.RequestTimeout)
		})
	})

	t.Run("Given the request timeout is negative", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"ipfs_check.base_url":                "https://ipfs-check.example.com",
			"ipfs_check.request_timeout_seconds": -1,
		}, "."), nil)

		t.Run("Should fall back to 90 seconds", func(t *testing.T) {
			t.Parallel()

			result := settings.NewIpfsCheckSettings(zap.NewNop(), k)

			assert.Equal(t, 90*time.Second, result.RequestTimeout)
		})
	})
}
