package subscribers

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionAgentImageProposalCreatedSubscriberInterface interface {
	SubscribeAgentImageProposalCreated(ctx context.Context, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1AgentImageProposalCreated) (ethereum.Subscription, error)
}
