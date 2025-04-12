package event_subscribers

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
)

type AgenticAIAgentCollectionMintedSubscriberInterface interface {
	SubscribeMinted(ctx context.Context, logs chan<- *abi.AgenticAIAgentCollectionMinted) (ethereum.Subscription, error)
}
