package interfaces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/ethereum/go-ethereum"
)

type CollectionCreatedSubscriptionProviderInterface interface {
	StartCollectionCreatedSubscription(
		ctx context.Context,
		chainId uint64,
	) (<-chan *abi.AgentCollectionsManagerCollectionCreated, ethereum.Subscription, error)
}
