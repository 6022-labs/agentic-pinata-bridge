package services

import (
	"math/big"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
)

type AgenticAIAgentCollectionRequester struct {
	client                           *ethclient.Client
	agenticAIAgentCollectionSettings *settings.AgenticAIAgentCollectionSettings
}

type newAgenticAIAgentCollectionRequesterParams struct {
	dig.In

	Client                           *ethclient.Client `name:"http"`
	AgenticAIAgentCollectionSettings *settings.AgenticAIAgentCollectionSettings
}

func NewAgenticAIAgentCollectionRequester(
	params newAgenticAIAgentCollectionRequesterParams,
) *AgenticAIAgentCollectionRequester {
	return &AgenticAIAgentCollectionRequester{
		client:                           params.Client,
		agenticAIAgentCollectionSettings: params.AgenticAIAgentCollectionSettings,
	}
}

func (a *AgenticAIAgentCollectionRequester) GetAllTokenIds() ([]big.Int, error) {
	agenticAIAgentCollection, err := abi.NewAgenticAIAgentCollection(
		a.agenticAIAgentCollectionSettings.SmartContractAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	nextTokenId, err := agenticAIAgentCollection.NextTokenId(nil)
	if err != nil {
		return nil, err
	}
	if nextTokenId.Cmp(big.NewInt(1)) == 0 {
		return []big.Int{}, nil
	}

	// From 1 to nextTokenId - 1
	tokenIds := make([]big.Int, 0, nextTokenId.Int64()-1)
	for i := 1; i < int(nextTokenId.Int64()); i++ {
		tokenId := big.NewInt(int64(i))
		tokenIds = append(tokenIds, *tokenId)
	}

	return tokenIds, nil
}

func (a *AgenticAIAgentCollectionRequester) GetAgentImage(agentTokenId big.Int) (*string, error) {
	agenticAIAgentCollection, err := abi.NewAgenticAIAgentCollection(
		a.agenticAIAgentCollectionSettings.SmartContractAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	image, err := agenticAIAgentCollection.ImageOf(nil, &agentTokenId)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
