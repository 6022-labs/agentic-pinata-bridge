package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWhenLoadingHostSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given no host configuration", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")

		t.Run("Should fall back to the defaults", func(t *testing.T) {
			t.Parallel()

			result := settings.NewHostSettings(zap.NewNop(), k)

			assert.True(t, result.UseApi)
			assert.True(t, result.UseListeners)
			assert.Equal(t, 3000, result.ApiPort)
		})
	})

	t.Run("Given an explicit host configuration", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"host.use_api":       false,
			"host.use_listeners": false,
			"host.api_port":      8080,
		}, "."), nil)

		t.Run("Should honor every provided value", func(t *testing.T) {
			t.Parallel()

			result := settings.NewHostSettings(zap.NewNop(), k)

			assert.False(t, result.UseApi)
			assert.False(t, result.UseListeners)
			assert.Equal(t, 8080, result.ApiPort)
		})
	})

	t.Run("Given a non-positive api port", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"host.api_port": 0,
		}, "."), nil)

		t.Run("Should warn and default the port to 3000", func(t *testing.T) {
			t.Parallel()

			result := settings.NewHostSettings(zap.NewNop(), k)

			assert.Equal(t, 3000, result.ApiPort)
		})
	})
}
