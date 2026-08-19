package interfaces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type MintedSubscriptionProviderInterface interface {
	StartMintedSubscription(
		ctx context.Context,
		chainId uint64,
		agentCollectionAddresses []common.Address,
	) (<-chan *abi.AgentCollectionV1Minted, ethereum.Subscription, error)
}
