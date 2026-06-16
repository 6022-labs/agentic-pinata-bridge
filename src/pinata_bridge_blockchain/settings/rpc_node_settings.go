package settings

import (
	"go.uber.org/zap"
)

type RpcNodeSettings struct {
	WsUrl   string
	HttpUrl string
}

func NewRpcNodeSettings(logger *zap.Logger, chains *ChainsSettings) *RpcNodeSettings {
	if chains == nil || len(*chains) != 1 {
		logger.Fatal("single-chain build expects exactly one configured chain in 'chains'")
	}

	for _, chain := range *chains {
		return &RpcNodeSettings{
			WsUrl:   chain.RpcWsUrl,
			HttpUrl: chain.RpcHttpUrl,
		}
	}

	// Unreachable: the length check above guarantees exactly one entry.
	return nil
}
