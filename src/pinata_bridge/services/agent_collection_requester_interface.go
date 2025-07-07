package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AgentCollectionRequesterInterface interface {
	GetAllTokenIds(collectionAddress common.Address) ([]big.Int, error)
	GetAgentImages(collectionAddress common.Address, agentTokenId big.Int) ([]string, error)
}
