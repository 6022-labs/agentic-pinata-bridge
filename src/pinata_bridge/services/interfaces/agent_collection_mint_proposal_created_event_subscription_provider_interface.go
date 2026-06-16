package interfaces

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface interface {
	StartMintProposalCreatedSubscription(chainId uint64, agentCollectionAddress common.Address) (<-chan *abi.AgentCollectionV1MintProposalCreated, ethereum.Subscription, error)
}
