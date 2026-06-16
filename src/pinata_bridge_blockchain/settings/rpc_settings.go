package settings

import (
	"strings"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const RpcSettingsKey = "chains"

type RpcConfig struct {
	HttpUrl string `koanf:"rpc_http_url"`
	WsUrl   string `koanf:"rpc_ws_url"`
}

type RpcSettings map[uint64]*RpcConfig

func NewRpcSettings(logger *zap.Logger, k *koanf.Koanf) *RpcSettings {
	settings := RpcSettings{}
	if err := k.UnmarshalWithConf(RpcSettingsKey, &settings, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
		logger.Fatal("failed to unmarshal rpc settings", zap.Error(err))
	}

	for chainId, rpc := range settings {
		if rpc == nil {
			logger.Fatal("rpc entry is nil", zap.Uint64("chain_id", chainId))
		}
		rpc.HttpUrl = strings.TrimSpace(rpc.HttpUrl)
		rpc.WsUrl = strings.TrimSpace(rpc.WsUrl)

		if rpc.HttpUrl == "" {
			logger.Fatal("chain rpc_http_url is required", zap.Uint64("chain_id", chainId))
		}
		if rpc.WsUrl == "" {
			logger.Fatal("chain rpc_ws_url is required", zap.Uint64("chain_id", chainId))
		}
	}

	return &settings
}

func (s *RpcSettings) Get(chainId uint64) *RpcConfig {
	if s == nil {
		return nil
	}
	return (*s)[chainId]
}
