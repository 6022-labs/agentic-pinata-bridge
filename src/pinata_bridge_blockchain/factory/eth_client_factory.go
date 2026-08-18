package factory

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// EthClientFactory dials and caches one HTTP and one WS client per chain for this adapter.
type EthClientFactory struct {
	rpcSettings          *settings.RpcSettings
	blockchainRpcMetrics metrics_interfaces.BlockchainRpcMetricsInterface
	mu                   sync.Mutex
	httpClients          map[uint64]*ethclient.Client
	wsClients            map[uint64]*ethclient.Client
}

func NewEthClientFactory(
	rpcSettings *settings.RpcSettings,
	blockchainRpcMetrics metrics_interfaces.BlockchainRpcMetricsInterface,
) *EthClientFactory {
	return &EthClientFactory{
		rpcSettings:          rpcSettings,
		blockchainRpcMetrics: blockchainRpcMetrics,
		httpClients:          map[uint64]*ethclient.Client{},
		wsClients:            map[uint64]*ethclient.Client{},
	}
}

// Http returns the HTTP client for a chain, dialing it on first use.
func (f *EthClientFactory) Http(chainId uint64) (*ethclient.Client, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if client, ok := f.httpClients[chainId]; ok {
		return client, nil
	}

	rpcConfig := f.rpcSettings.Get(chainId)
	if rpcConfig == nil {
		return nil, fmt.Errorf("no rpc configured for chain %d", chainId)
	}

	httpClient := &http.Client{
		Transport: metrics.NewRpcRoundTripper(
			http.DefaultTransport,
			f.blockchainRpcMetrics,
			strconv.FormatUint(chainId, 10),
		),
	}

	rpcClient, err := rpc.DialOptions(context.Background(), rpcConfig.HttpUrl, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("dial http rpc for chain %d: %w", chainId, err)
	}

	client := ethclient.NewClient(rpcClient)
	f.httpClients[chainId] = client
	return client, nil
}

// Ws returns the WS client for a chain, dialing it on first use; a subscription is not a
// request/response the rpc round tripper can measure, so this one stays uninstrumented.
func (f *EthClientFactory) Ws(chainId uint64) (*ethclient.Client, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if client, ok := f.wsClients[chainId]; ok {
		return client, nil
	}

	rpcConfig := f.rpcSettings.Get(chainId)
	if rpcConfig == nil {
		return nil, fmt.Errorf("no rpc configured for chain %d", chainId)
	}

	client, err := ethclient.Dial(rpcConfig.WsUrl)
	if err != nil {
		return nil, fmt.Errorf("dial ws rpc for chain %d: %w", chainId, err)
	}

	f.wsClients[chainId] = client
	return client, nil
}
