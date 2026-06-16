package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func panicLogger() *zap.Logger {
	return zap.NewNop().WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic))
}

func TestWhenLoadingPinataSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given a fully populated pinata configuration", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"pinata.api_key":  "  secret-key  ",
			"pinata.base_url": "  https://api.pinata.cloud  ",
		}, "."), nil)

		t.Run("Should load and trim every field", func(t *testing.T) {
			t.Parallel()

			result := settings.NewPinataSettings(zap.NewNop(), k)

			assert.Equal(t, "secret-key", result.ApiKey)
			assert.Equal(t, "https://api.pinata.cloud", result.BaseUrl)
		})
	})

	t.Run("Given the api key is missing", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"pinata.api_key":  "   ",
			"pinata.base_url": "https://api.pinata.cloud",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewPinataSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given the base url is missing", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"pinata.api_key":  "secret-key",
			"pinata.base_url": "",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewPinataSettings(panicLogger(), k)
			})
		})
	})
}
