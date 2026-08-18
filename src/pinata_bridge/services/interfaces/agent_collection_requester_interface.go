package interfaces

import (
	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionRequesterInterface interface {
	GetAllTokenIds(ctx context.Context, chainId uint64, collectionAddress common.Address) ([]big.Int, error)
	GetAgentImages(
		ctx context.Context,
		chainId uint64,
		collectionAddress common.Address,
		agentTokenId big.Int,
	) ([]string, error)
	GetMintProposalImages(
		ctx context.Context,
		chainId uint64,
		collectionAddress common.Address,
		proposalId big.Int,
	) ([]string, error)
	GetAgentImageProposalImage(
		ctx context.Context,
		chainId uint64,
		collectionAddress common.Address,
		proposalId big.Int,
	) (*string, error)
}
