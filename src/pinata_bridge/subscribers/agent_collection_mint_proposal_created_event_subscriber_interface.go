package subscribers

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionMintProposalCreatedSubscriberInterface interface {
	SubscribeMintProposalCreated(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1MintProposalCreated) (ethereum.Subscription, error)
}
