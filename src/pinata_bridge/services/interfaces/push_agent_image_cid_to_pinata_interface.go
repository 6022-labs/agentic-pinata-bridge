package interfaces

import (
	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PushAgentImageCidToPinataInterface interface {
	PushMissingImageCids(ctx context.Context) error
	PushMissingImagesOfAgent(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, agentCollectionTokenId big.Int) error
	PushImagesOfMintProposal(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
	PushImageOfAgentImageProposal(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
	PushFromCid(ctx context.Context, cid string) error
}
