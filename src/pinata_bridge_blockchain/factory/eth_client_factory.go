package factory

import (
	"fmt"
	"sync"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClientFactory dials and caches one HTTP and one WS client per chain.
type EthClientFactory struct {
	rpcSettings *settings.RpcSettings
	mu          sync.Mutex
	httpClients map[uint64]*ethclient.Client
	wsClients   map[uint64]*ethclient.Client
}

func NewEthClientFactory(rpcSettings *settings.RpcSettings) *EthClientFactory {
	return &EthClientFactory{
		rpcSettings: rpcSettings,
		httpClients: map[uint64]*ethclient.Client{},
		wsClients:   map[uint64]*ethclient.Client{},
	}
}

// Http returns the HTTP client for a chain, dialing it on first use.
func (f *EthClientFactory) Http(chainId uint64) (*ethclient.Client, error) {
	return f.dial(chainId, f.httpClients, func(rpc *settings.RpcConfig) string { return rpc.HttpUrl }, "http")
}

// Ws returns the WS client for a chain, dialing it on first use.
func (f *EthClientFactory) Ws(chainId uint64) (*ethclient.Client, error) {
	return f.dial(chainId, f.wsClients, func(rpc *settings.RpcConfig) string { return rpc.WsUrl }, "ws")
}

func (f *EthClientFactory) dial(
	chainId uint64,
	cache map[uint64]*ethclient.Client,
	url func(*settings.RpcConfig) string,
	kind string,
) (*ethclient.Client, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if client, ok := cache[chainId]; ok {
		return client, nil
	}

	rpc := f.rpcSettings.Get(chainId)
	if rpc == nil {
		return nil, fmt.Errorf("no rpc configured for chain %d", chainId)
	}

	client, err := ethclient.Dial(url(rpc))
	if err != nil {
		return nil, fmt.Errorf("dial %s rpc for chain %d: %w", kind, chainId, err)
	}
	cache[chainId] = client
	return client, nil
}
