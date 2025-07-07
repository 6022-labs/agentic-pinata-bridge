package subscribers

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
)

type AgentCollectionsManagerCollectionCreatedSubscriberInterface interface {
	SubscribeCollectionCreated(ctx context.Context, logs chan<- *abi.AgentCollectionsManagerCollectionCreated) (ethereum.Subscription, error)
}
