package interfaces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type AgentImageProposalCreatedSubscriptionProviderInterface interface {
	StartAgentImageProposalCreatedSubscription(
		ctx context.Context,
		chainId uint64,
		agentCollectionAddress common.Address,
	) (<-chan *abi.AgentCollectionV1AgentImageProposalCreated, ethereum.Subscription, error)
}
