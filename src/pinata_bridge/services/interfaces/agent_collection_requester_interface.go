package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionRequesterInterface interface {
	GetAllTokenIds(chainId uint64, collectionAddress common.Address) ([]big.Int, error)
	GetAgentImages(chainId uint64, collectionAddress common.Address, agentTokenId big.Int) ([]string, error)
	GetMintProposalImages(chainId uint64, collectionAddress common.Address, proposalId big.Int) ([]string, error)
	GetAgentImageProposalImage(chainId uint64, collectionAddress common.Address, proposalId big.Int) (*string, error)
}
