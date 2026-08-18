package interfaces

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionsManagerRequesterInterface interface {
	GetAllCollectionAddresses(ctx context.Context, chainId uint64) ([]common.Address, error)
}
