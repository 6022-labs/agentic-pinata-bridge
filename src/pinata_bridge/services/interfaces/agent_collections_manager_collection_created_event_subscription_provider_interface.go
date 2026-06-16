package interfaces

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
)

type AgentCollectionsManagerCollectionCreatedEventSubscriptionProviderInterface interface {
	StartCollectionCreatedSubscription(chainId uint64) (<-chan *abi.AgentCollectionsManagerCollectionCreated, ethereum.Subscription, error)
}
