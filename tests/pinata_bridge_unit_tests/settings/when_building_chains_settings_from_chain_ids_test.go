package settings_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/stretchr/testify/assert"
)

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
