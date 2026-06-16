package settings

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const AgentCollectionsManagersSettingsKey = "agent_collections_managers"

type AgentCollectionsManagersSettings map[uint64]common.Address

func NewAgentCollectionsManagersSettings(logger *zap.Logger, k *koanf.Koanf) *AgentCollectionsManagersSettings {
	raw := map[uint64]string{}
	if err := k.UnmarshalWithConf(AgentCollectionsManagersSettingsKey, &raw, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
		logger.Fatal("failed to unmarshal agent_collections_managers settings", zap.Error(err))
	}

	settings := AgentCollectionsManagersSettings{}
	for chainId, addr := range raw {
		addr = strings.TrimSpace(addr)
		if addr == "" {
			logger.Fatal("agent_collections_manager address is required", zap.Uint64("chain_id", chainId))
		}
		if !common.IsHexAddress(addr) {
			logger.Fatal(
				"agent_collections_manager address is not a valid hex address",
				zap.Uint64("chain_id", chainId),
				zap.String("address", addr),
			)
		}
		settings[chainId] = common.HexToAddress(addr)
	}

	if len(settings) == 0 {
		logger.Fatal("agent_collections_managers settings is empty: configure at least one via appsettings.json or AGENT_COLLECTIONS_MANAGERS__<chain_id> env vars")
	}

	return &settings
}

func (s *AgentCollectionsManagersSettings) Get(chainId uint64) (common.Address, bool) {
	if s == nil {
		return common.Address{}, false
	}
	addr, ok := (*s)[chainId]
	return addr, ok
}
