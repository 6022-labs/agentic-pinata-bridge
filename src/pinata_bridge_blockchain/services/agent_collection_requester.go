package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	pinata_bridge_abi "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/samber/lo"
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
	AgentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
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
	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
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

func (a *AgentCollectionRequester) GetMintProposalImages(collectionAddress common.Address, proposalId big.Int) ([]string, error) {
	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	mintProposals, err := a.getMintProposals(agentCollection, collectionAddress)
	if err != nil {
		return nil, err
	}

	for _, proposal := range mintProposals {
		if proposal.Id.Cmp(&proposalId) == 0 {
			return lo.Map(proposal.Images, func(image pinata_bridge_abi.KeyValue, _ int) string {
				return image.Value
			}), nil
		}
	}

	return nil, fmt.Errorf("mint proposal with ID %s not found in collection %s", proposalId.String(), collectionAddress.Hex())
}

func (a *AgentCollectionRequester) GetAgentImageProposalImage(collectionAddress common.Address, proposalId big.Int) (*string, error) {
	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		a.client,
	)
	if err != nil {
		return nil, err
	}

	addOrUpdateImageProposals, err := a.getAddOrUpdateImageProposals(agentCollection, collectionAddress)
	if err != nil {
		return nil, err
	}

	for _, proposal := range addOrUpdateImageProposals {
		if proposal.Id.Cmp(&proposalId) == 0 {
			return &proposal.Image.Value, nil
		}
	}

	return nil, fmt.Errorf("add or update image proposal with ID %s not found in collection %s", proposalId.String(), collectionAddress.Hex())
}

func (a *AgentCollectionRequester) getMintProposals(agentCollection *pinata_bridge_abi.AgentCollectionV1, agentCollectionAddress common.Address) ([]pinata_bridge_abi.MintProposal, error) {
	mintProposalLength, err := agentCollection.MintProposalsLength(nil)
	if err != nil {
		return nil, err
	}

	agentCollectionABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionV1ABI)) // You must expose raw ABI string
	if err != nil {
		return nil, err
	}

	batch := make([]rpc.BatchElem, 0, mintProposalLength.Int64())
	results := make([]json.RawMessage, mintProposalLength.Int64())

	for i := uint64(0); i < mintProposalLength.Uint64(); i++ {
		data, err := agentCollectionABI.Pack("mintProposal", big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		msg := map[string]interface{}{
			"to":   agentCollectionAddress.Hex(),
			"data": "0x" + common.Bytes2Hex(data),
		}

		batch = append(batch, rpc.BatchElem{
			Method: "eth_call",
			Args:   []interface{}{msg, "latest"},
			Result: &results[i],
		})
	}

	if err := a.client.Client().BatchCallContext(context.Background(), batch); err != nil {
		return nil, err
	}

	mintProposals := make([]pinata_bridge_abi.MintProposal, 0, mintProposalLength.Int64())
	for _, result := range results {
		if result == nil {
			continue
		}

		var mintProposal pinata_bridge_abi.MintProposal
		if err := agentCollectionABI.UnpackIntoInterface(&mintProposal, "mintProposal", result); err != nil {
			return nil, err
		}

		mintProposals = append(mintProposals, mintProposal)
	}

	return mintProposals, nil
}

func (a *AgentCollectionRequester) getAddOrUpdateImageProposals(agentCollection *pinata_bridge_abi.AgentCollectionV1, agentCollectionAddress common.Address) ([]pinata_bridge_abi.AddOrUpdateImageProposal, error) {
	addOrUpdateImageProposalLength, err := agentCollection.AddOrUpdateImageProposalsLength(nil)
	if err != nil {
		return nil, err
	}

	agentCollectionABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionV1ABI)) // You must expose raw ABI string
	if err != nil {
		return nil, err
	}

	batch := make([]rpc.BatchElem, 0, addOrUpdateImageProposalLength.Int64())
	results := make([]json.RawMessage, addOrUpdateImageProposalLength.Int64())

	for i := uint64(0); i < addOrUpdateImageProposalLength.Uint64(); i++ {
		data, err := agentCollectionABI.Pack("addOrUpdateImageProposal", big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		msg := map[string]interface{}{
			"to":   agentCollectionAddress.Hex(),
			"data": "0x" + common.Bytes2Hex(data),
		}

		batch = append(batch, rpc.BatchElem{
			Method: "eth_call",
			Args:   []interface{}{msg, "latest"},
			Result: &results[i],
		})
	}

	if err := a.client.Client().BatchCallContext(context.Background(), batch); err != nil {
		return nil, err
	}

	addOrUpdateImageProposals := make([]pinata_bridge_abi.AddOrUpdateImageProposal, 0, addOrUpdateImageProposalLength.Int64())

	for _, result := range results {
		if result == nil {
			continue
		}

		var addOrUpdateImageProposal pinata_bridge_abi.AddOrUpdateImageProposal
		if err := agentCollectionABI.UnpackIntoInterface(&addOrUpdateImageProposal, "addOrUpdateImageProposal", result); err != nil {
			return nil, err
		}

		addOrUpdateImageProposals = append(addOrUpdateImageProposals, addOrUpdateImageProposal)
	}

	return addOrUpdateImageProposals, nil
}
