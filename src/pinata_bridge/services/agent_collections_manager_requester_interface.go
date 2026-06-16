package services

import "github.com/ethereum/go-ethereum/common"

type AgentCollectionsManagerRequesterInterface interface {
	GetAllCollectionAddresses(chainId uint64) ([]common.Address, error)
}
