package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func panicLogger() *zap.Logger {
	return zap.NewNop().WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic))
}

func TestWhenLoadingChainsSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given several chains in unsorted order", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"chains.137.rpc_http_url":   "https://polygon",
			"chains.80002.rpc_http_url": "https://amoy",
			"chains.1.rpc_http_url":     "https://mainnet",
		}, "."), nil)

		t.Run("Should expose chain ids sorted ascending and answer Has", func(t *testing.T) {
			t.Parallel()

			result := settings.NewChainsSettings(zap.NewNop(), k)

			assert.Equal(t, []uint64{1, 137, 80002}, result.ChainIds())
			assert.True(t, result.Has(80002))
			assert.False(t, result.Has(999))
		})
	})

	t.Run("Given no chains are configured", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewChainsSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given a non-numeric chain id", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"chains.not-a-number.rpc_http_url": "https://example",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewChainsSettings(panicLogger(), k)
			})
		})
	})
}

func TestWhenBuildingChainsSettingsFromChainIds(t *testing.T) {
	t.Parallel()

	t.Run("Given an unsorted list of chain ids", func(t *testing.T) {
		t.Parallel()

		t.Run("Should sort them without mutating the input slice", func(t *testing.T) {
			t.Parallel()

			input := []uint64{80002, 1, 137}
			result := settings.NewChainsSettingsFromChainIds(input)

			assert.Equal(t, []uint64{1, 137, 80002}, result.ChainIds())
			assert.Equal(t, []uint64{80002, 1, 137}, input)
		})
	})
}
