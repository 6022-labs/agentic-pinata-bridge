package settings

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

type RpcNodeSettings struct {
	WsUrl   string
	HttpUrl string
}

func NewRpcNodeSettings(logger *zap.Logger) *RpcNodeSettings {
	wsUrl := os.Getenv("RPC_NODE_WS_URL")
	if len(strings.TrimSpace(wsUrl)) == 0 {
		logger.Fatal("please set your RPC_NODE_WS_URL value in your environment")
	}

	httpUrl := os.Getenv("RPC_NODE_HTTP_URL")
	if len(strings.TrimSpace(httpUrl)) == 0 {
		logger.Fatal("please set your RPC_NODE_HTTP_URL value in your environment")
	}

	return &RpcNodeSettings{
		WsUrl:   wsUrl,
		HttpUrl: httpUrl,
	}
}
