package interfaces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionAgentImageProposalCreatedSubscriberInterface interface {
	SubscribeAgentImageProposalCreated(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1AgentImageProposalCreated) (ethereum.Subscription, error)
}
