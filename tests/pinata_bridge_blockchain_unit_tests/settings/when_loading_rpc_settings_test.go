package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func panicLogger() *zap.Logger {
	return zap.NewNop().WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic))
}

func TestWhenLoadingRpcSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given a fully populated chain entry", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"chains.80002.rpc_http_url": "  https://amoy.example/http  ",
			"chains.80002.rpc_ws_url":   "  wss://amoy.example/ws  ",
		}, "."), nil)

		t.Run("Should load and trim the urls and resolve them via Get", func(t *testing.T) {
			t.Parallel()

			result := settings.NewRpcSettings(zap.NewNop(), k)

			cfg := result.Get(80002)
			assert.NotNil(t, cfg)
			assert.Equal(t, "https://amoy.example/http", cfg.HttpUrl)
			assert.Equal(t, "wss://amoy.example/ws", cfg.WsUrl)
			assert.Nil(t, result.Get(999))
		})
	})

	t.Run("Given a chain entry missing its http url", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"chains.80002.rpc_http_url": "   ",
			"chains.80002.rpc_ws_url":   "wss://amoy.example/ws",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewRpcSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given a chain entry missing its ws url", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"chains.80002.rpc_http_url": "https://amoy.example/http",
			"chains.80002.rpc_ws_url":   "",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewRpcSettings(panicLogger(), k)
			})
		})
	})
}
