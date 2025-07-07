package services

import "github.com/ethereum/go-ethereum/common"

type AgentCollectionsManagerRequesterInterface interface {
	GetAllCollectionAddresses() ([]common.Address, error)
}
