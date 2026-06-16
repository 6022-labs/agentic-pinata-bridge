package settings

import (
	"slices"
	"strconv"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const ChainsSettingsKey = "chains"

type ChainsSettings struct {
	chainIds []uint64
}

func NewChainsSettings(logger *zap.Logger, k *koanf.Koanf) *ChainsSettings {
	keys := k.MapKeys(ChainsSettingsKey)
	if len(keys) == 0 {
		logger.Fatal("chains settings is empty: configure at least one chain via appsettings.json or CHAINS__<chain_id>__* env vars")
	}

	chainIds := make([]uint64, 0, len(keys))
	for _, key := range keys {
		chainId, err := strconv.ParseUint(key, 10, 64)
		if err != nil {
			logger.Fatal("invalid chain id in chains settings", zap.String("chain_id", key), zap.Error(err))
		}
		chainIds = append(chainIds, chainId)
	}
	slices.Sort(chainIds)

	return &ChainsSettings{chainIds: chainIds}
}

func NewChainsSettingsFromChainIds(chainIds []uint64) *ChainsSettings {
	sorted := slices.Clone(chainIds)
	slices.Sort(sorted)
	return &ChainsSettings{chainIds: sorted}
}

func (s *ChainsSettings) ChainIds() []uint64 {
	return s.chainIds
}

func (s *ChainsSettings) Has(chainId uint64) bool {
	return slices.Contains(s.chainIds, chainId)
}
