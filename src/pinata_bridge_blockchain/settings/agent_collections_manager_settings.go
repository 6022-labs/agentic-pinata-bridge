package settings

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionsManagerSettings struct {
	SmartContractAddress common.Address
}

func NewAgentCollectionsManagerSettings(logger *zap.Logger) *AgentCollectionsManagerSettings {
	smartContractAddressStr := os.Getenv("AGENT_COLLECTIONS_MANAGER_SMART_CONTRACT_ADDRESS")
	if len(strings.TrimSpace(smartContractAddressStr)) == 0 {
		logger.Fatal("please set your AGENT_COLLECTIONS_MANAGER_SMART_CONTRACT_ADDRESS value in your environment")
	}

	if !common.IsHexAddress(smartContractAddressStr) {
		logger.Fatal("AGENT_COLLECTIONS_MANAGER_SMART_CONTRACT_ADDRESS is not a valid Ethereum address")
	}

	smartContractAddress := common.HexToAddress(smartContractAddressStr)

	return &AgentCollectionsManagerSettings{
		SmartContractAddress: smartContractAddress,
	}
}
