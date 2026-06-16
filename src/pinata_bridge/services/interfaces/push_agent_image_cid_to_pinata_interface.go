package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PushAgentImageCidToPinataInterface interface {
	PushMissingImageCids() error
	PushMissingImagesOfAgent(chainId uint64, agentCollectionAddress common.Address, agentCollectionTokenId big.Int) error
	PushImagesOfMintProposal(chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
	PushImageOfAgentImageProposal(chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
	PushFromCid(cid string) error
}
