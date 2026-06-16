package settings

import (
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionsManagerSettings struct {
	SmartContractAddress common.Address
}

// NewAgentCollectionsManagerSettings derives the single-chain manager address from the
// per-chain managers map. This is a temporary single-chain adapter: the multi-chain
// follow-up removes it and lets services consume AgentCollectionsManagersSettings directly.
func NewAgentCollectionsManagerSettings(
	logger *zap.Logger,
	chains *ChainsSettings,
	managers *AgentCollectionsManagersSettings,
) *AgentCollectionsManagerSettings {
	if chains == nil || len(*chains) != 1 {
		logger.Fatal("single-chain build expects exactly one configured chain in 'chains'")
	}

	for chainId := range *chains {
		address, ok := managers.Get(chainId)
		if !ok {
			logger.Fatal(
				"no agent_collections_manager address configured for the chain",
				zap.Uint64("chain_id", chainId),
			)
		}

		return &AgentCollectionsManagerSettings{
			SmartContractAddress: address,
		}
	}

	// Unreachable: the length check above guarantees exactly one entry.
	return nil
}
