package services

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	pinata_bridge_abi "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
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

		msg := map[string]interface{}{
			"to":   a.agentCollectionsManagerSettings.SmartContractAddress.Hex(),
			"data": "0x" + common.Bytes2Hex(data),
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
		var hexStr string
		if err := json.Unmarshal(res, &hexStr); err != nil {
			return nil, err
		}
		hexStr = strings.TrimPrefix(hexStr, "0x")
		dataBytes := common.FromHex("0x" + hexStr)
		var unpacked common.Address
		if err := agentABI.UnpackIntoInterface(&unpacked, "collections", dataBytes); err != nil {
			return nil, err
		}
		addresses = append(addresses, unpacked)
	}

	return addresses, nil
}
