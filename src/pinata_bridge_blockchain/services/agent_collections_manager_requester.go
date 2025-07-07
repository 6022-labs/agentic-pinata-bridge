package services

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	pinata_bridge_abi "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/dig"
)

type AgentCollectionsManagerRequester struct {
	client                          *ethclient.Client
	agentCollectionsManagerSettings *settings.AgentCollectionsManagerSettings
}

type newAgentCollectionsManagerRequesterParams struct {
	dig.In

	Client                          *ethclient.Client `name:"http"`
	AgentCollectionsManagerSettings *settings.AgentCollectionsManagerSettings
}

func NewAgentCollectionsManagerRequester(
	params newAgentCollectionsManagerRequesterParams,
) *AgentCollectionsManagerRequester {
	return &AgentCollectionsManagerRequester{
		client:                          params.Client,
		agentCollectionsManagerSettings: params.AgentCollectionsManagerSettings,
	}
}

func (a *AgentCollectionsManagerRequester) GetAllCollectionAddresses() ([]common.Address, error) {
	agentABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionsManagerABI)) // You must expose raw ABI string
	if err != nil {
		return nil, err
	}

	contract, err := pinata_bridge_abi.NewAgentCollectionsManager(
		a.agentCollectionsManagerSettings.SmartContractAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	nextId, err := contract.NextCollectionId(nil)
	if err != nil {
		return nil, err
	}

	collectionCount := int(nextId.Int64()) - 1
	if collectionCount <= 0 {
		return nil, nil
	}

	batch := make([]rpc.BatchElem, 0, collectionCount)
	results := make([]json.RawMessage, collectionCount)

	for i := 1; i <= collectionCount; i++ {
		data, err := agentABI.Pack("collections", big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		msg := ethereum.CallMsg{
			To:   &a.agentCollectionsManagerSettings.SmartContractAddress,
			Data: data,
		}

		batch = append(batch, rpc.BatchElem{
			Method: "eth_call",
			Args:   []interface{}{msg, "latest"},
			Result: &results[i-1],
		})
	}

	if err := a.client.Client().BatchCallContext(context.Background(), batch); err != nil {
		return nil, err
	}

	addresses := make([]common.Address, 0, collectionCount)
	for _, res := range results {
		var unpacked []interface{}
		if err := agentABI.UnpackIntoInterface(&unpacked, "collections", res); err != nil {
			return nil, err
		}
		addresses = append(addresses, unpacked[0].(common.Address))
	}

	return addresses, nil
}
