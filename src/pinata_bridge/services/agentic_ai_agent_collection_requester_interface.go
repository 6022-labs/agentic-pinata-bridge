package services

import (
	"math/big"
)

type AgenticAIAgentCollectionRequesterInterface interface {
	GetAgentImage(agentTokenId big.Int) (*string, error)
}
