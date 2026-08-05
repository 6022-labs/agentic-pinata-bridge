package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWhenLoadingAgentCollectionsManagersSettings(t *testing.T) {
	t.Parallel()

	t.Run("Given the embedded agent_collections_managers.json", func(t *testing.T) {
		t.Parallel()

		result := settings.NewAgentCollectionsManagersSettings(zap.NewNop())

		t.Run("Should expose the deployed manager address for polygon amoy", func(t *testing.T) {
			t.Parallel()

			addr, ok := result.Get(80002)
			assert.True(t, ok)
			assert.Equal(t, common.HexToAddress("0x8095A0D4bE42A61db9e609aA693f516044Bda990"), addr)
		})

		t.Run("Should expose the deployed manager address for polygon mainnet", func(t *testing.T) {
			t.Parallel()

			addr, ok := result.Get(137)
			assert.True(t, ok)
			assert.Equal(t, common.HexToAddress("0xb0dc8a83c700A9BBcc53cA1a2C6993a63129d2F6"), addr)
		})

		t.Run("Should report not-ok for an unconfigured chain", func(t *testing.T) {
			t.Parallel()

			_, ok := result.Get(999)
			assert.False(t, ok)
		})
	})
}
