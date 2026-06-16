package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const validManagerAddress = "0x52908400098527886E0F7030069857D2E4169EE7"

func TestWhenLoadingAgentCollectionsManagersSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid manager address per chain", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"agent_collections_managers.80002": "  " + validManagerAddress + "  ",
		}, "."), nil)

		t.Run("Should parse the address and resolve it via Get", func(t *testing.T) {
			t.Parallel()

			result := settings.NewAgentCollectionsManagersSettings(zap.NewNop(), k)

			addr, ok := result.Get(80002)
			assert.True(t, ok)
			assert.Equal(t, common.HexToAddress(validManagerAddress), addr)

			_, missing := result.Get(999)
			assert.False(t, missing)
		})
	})

	t.Run("Given an empty manager address", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"agent_collections_managers.80002": "   ",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewAgentCollectionsManagersSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given an invalid hex address", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")
		_ = k.Load(confmap.Provider(map[string]any{
			"agent_collections_managers.80002": "not-an-address",
		}, "."), nil)

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewAgentCollectionsManagersSettings(panicLogger(), k)
			})
		})
	})

	t.Run("Given no managers are configured", func(t *testing.T) {
		t.Parallel()

		k := koanf.New(".")

		t.Run("Should fatal", func(t *testing.T) {
			t.Parallel()

			assert.Panics(t, func() {
				settings.NewAgentCollectionsManagersSettings(panicLogger(), k)
			})
		})
	})
}
