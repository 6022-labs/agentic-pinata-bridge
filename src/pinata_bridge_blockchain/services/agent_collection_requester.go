package services

import (
	"math/big"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
)

type AgentCollectionRequester struct {
	client *ethclient.Client
}

type newAgentCollectionRequesterParams struct {
	dig.In

	Client *ethclient.Client `name:"http"`
}

func NewAgentCollectionRequester(
	params newAgentCollectionRequesterParams,
) *AgentCollectionRequester {
	return &AgentCollectionRequester{
		client: params.Client,
	}
}

func (a *AgentCollectionRequester) GetAllTokenIds(collectionAddress common.Address) ([]big.Int, error) {
	AgentCollection, err := abi.NewAgentCollectionV1(
		collectionAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	nextTokenId, err := AgentCollection.NextTokenId(nil)
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

func (a *AgentCollectionRequester) GetAgentImages(collectionAddress common.Address, agentTokenId big.Int) ([]string, error) {
	agentCollection, err := abi.NewAgentCollectionV1(
		collectionAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	imagesKeyValues, err := agentCollection.ImagesOf(nil, &agentTokenId)
	if err != nil {
		return nil, err
	}

	images := make([]string, 0, len(imagesKeyValues))
	for _, kv := range imagesKeyValues {
		images = append(images, kv.Value)
	}

	return images, nil
}
