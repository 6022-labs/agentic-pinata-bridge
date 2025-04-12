package services

import (
	"math/big"
)

type AgenticAIAgentCollectionRequesterInterface interface {
	GetAllTokenIds() ([]big.Int, error)
	GetAgentImage(agentTokenId big.Int) (*string, error)
}
