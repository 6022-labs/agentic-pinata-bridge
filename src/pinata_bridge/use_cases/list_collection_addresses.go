package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/ethereum/go-ethereum/common"
)

// ListCollectionAddresses returns every agent collection the manager governs on one chain.
type ListCollectionAddresses struct {
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface
}

func NewListCollectionAddresses(
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
) *ListCollectionAddresses {
	return &ListCollectionAddresses{agentCollectionsManagerRequester: agentCollectionsManagerRequester}
}

func (u *ListCollectionAddresses) Execute(ctx context.Context, chainId uint64) ([]common.Address, error) {
	return u.agentCollectionsManagerRequester.GetAllCollectionAddresses(ctx, chainId)
}
