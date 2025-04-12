package settings

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgenticAIAgentCollectionSettings struct {
	SmartContractAddress common.Address
}

func NewAgenticAIAgentCollectionSettings(logger *zap.Logger) *AgenticAIAgentCollectionSettings {
	smartContractAddressStr := os.Getenv("AGENTIC_AI_AGENT_COLLECTION_SMART_CONTRACT_ADDRESS")
	if len(strings.TrimSpace(smartContractAddressStr)) == 0 {
		logger.Fatal("please set your AGENTIC_AI_AGENT_COLLECTION_SMART_CONTRACT_ADDRESS value in your environment")
	}

	smartContractAddress := common.HexToAddress(smartContractAddressStr)

	return &AgenticAIAgentCollectionSettings{
		SmartContractAddress: smartContractAddress,
	}
}
