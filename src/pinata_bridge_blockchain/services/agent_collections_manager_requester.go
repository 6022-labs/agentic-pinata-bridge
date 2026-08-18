package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	pinata_bridge_abi "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type AgentCollectionsManagerRequester struct {
	ethClientFactory *factory.EthClientFactory
	managers         *settings.AgentCollectionsManagersSettings
}

func NewAgentCollectionsManagerRequester(
	ethClientFactory *factory.EthClientFactory,
	managers *settings.AgentCollectionsManagersSettings,
) *AgentCollectionsManagerRequester {
	return &AgentCollectionsManagerRequester{
		ethClientFactory: ethClientFactory,
		managers:         managers,
	}
}

func (a *AgentCollectionsManagerRequester) GetAllCollectionAddresses(ctx context.Context, chainId uint64) ([]common.Address, error) {
	client, err := a.ethClientFactory.Http(chainId)
	if err != nil {
		return nil, err
	}

	managerAddress, ok := a.managers.Get(chainId)
	if !ok {
		return nil, fmt.Errorf("no agent collections manager configured for chain %d", chainId)
	}

	agentABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionsManagerABI)) // You must expose raw ABI string
	if err != nil {
		return nil, err
	}

	contract, err := pinata_bridge_abi.NewAgentCollectionsManager(
		managerAddress,
		client,
	)
	if err != nil {
		return nil, err
	}

	nextId, err := contract.NextCollectionId(&bind.CallOpts{Context: ctx})
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
			"to":   managerAddress.Hex(),
			"data": "0x" + common.Bytes2Hex(data),
		}

		batch = append(batch, rpc.BatchElem{
			Method: "eth_call",
			Args:   []interface{}{msg, "latest"},
			Result: &results[i-1],
		})
	}

	if err := client.Client().BatchCallContext(ctx, batch); err != nil {
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
