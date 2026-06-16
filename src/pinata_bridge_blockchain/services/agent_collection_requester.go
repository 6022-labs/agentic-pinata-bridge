package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	pinata_bridge_abi "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/samber/lo"
)

type AgentCollectionRequester struct {
	ethClientFactory *factory.EthClientFactory
}

func NewAgentCollectionRequester(
	ethClientFactory *factory.EthClientFactory,
) *AgentCollectionRequester {
	return &AgentCollectionRequester{
		ethClientFactory: ethClientFactory,
	}
}

func (a *AgentCollectionRequester) GetAllTokenIds(chainId uint64, collectionAddress common.Address) ([]big.Int, error) {
	client, err := a.ethClientFactory.Http(chainId)
	if err != nil {
		return nil, err
	}

	AgentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create agent collection instance: %w", err)
	}

	nextTokenId, err := AgentCollection.NextTokenId(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get next token ID: %w", err)
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

func (a *AgentCollectionRequester) GetAgentImages(chainId uint64, collectionAddress common.Address, agentTokenId big.Int) ([]string, error) {
	client, err := a.ethClientFactory.Http(chainId)
	if err != nil {
		return nil, err
	}

	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create agent collection instance: %w", err)
	}

	imageCount, err := agentCollection.ImagesCountOf(nil, &agentTokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to get agent image count: %w", err)
	}

	imagesKeyValues := make([]pinata_bridge_abi.KeyValue, imageCount.Uint64())

	for i := uint64(0); i < imageCount.Uint64(); i++ {
		imageKeyValue, err := agentCollection.ImageOfByIndex(nil, &agentTokenId, big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to get agent image at index %d: %w", i, err)
		}
		imagesKeyValues[i] = imageKeyValue
	}

	images := make([]string, 0, len(imagesKeyValues))
	for _, kv := range imagesKeyValues {
		images = append(images, kv.Value)
	}

	return images, nil
}

func (a *AgentCollectionRequester) GetMintProposalImages(chainId uint64, collectionAddress common.Address, proposalId big.Int) ([]string, error) {
	client, err := a.ethClientFactory.Http(chainId)
	if err != nil {
		return nil, err
	}

	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create agent collection instance: %w", err)
	}

	mintProposals, err := a.getMintProposals(client, agentCollection, collectionAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get mint proposals: %w", err)
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

func (a *AgentCollectionRequester) GetAgentImageProposalImage(chainId uint64, collectionAddress common.Address, proposalId big.Int) (*string, error) {
	client, err := a.ethClientFactory.Http(chainId)
	if err != nil {
		return nil, err
	}

	agentCollection, err := pinata_bridge_abi.NewAgentCollectionV1(
		collectionAddress,
		client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create agent collection: %w", err)
	}

	addOrUpdateAgentImageProposals, err := a.getAddOrUpdateAgentImageProposals(client, agentCollection, collectionAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get add or update agent image proposals: %w", err)
	}

	for _, proposal := range addOrUpdateAgentImageProposals {
		if proposal.Id.Cmp(&proposalId) == 0 {
			return &proposal.Image.Value, nil
		}
	}

	return nil, fmt.Errorf("add or update agent image proposal with ID %s not found in collection %s", proposalId.String(), collectionAddress.Hex())
}

func (a *AgentCollectionRequester) getMintProposals(client *ethclient.Client, agentCollection *pinata_bridge_abi.AgentCollectionV1, agentCollectionAddress common.Address) ([]pinata_bridge_abi.MintProposal, error) {
	mintProposalLength, err := agentCollection.MintProposalsLength(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get mint proposal length: %w", err)
	}

	agentCollectionABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionV1ABI)) // You must expose raw ABI string
	if err != nil {
		return nil, fmt.Errorf("failed to parse agent collection ABI: %w", err)
	}

	batch := make([]rpc.BatchElem, 0, mintProposalLength.Int64())
	results := make([]hexutil.Bytes, mintProposalLength.Int64())

	for i := uint64(0); i < mintProposalLength.Uint64(); i++ {
		data, err := agentCollectionABI.Pack("mintProposal", big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to pack mintProposal call: %w", err)
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

	if err := client.Client().BatchCall(batch); err != nil {
		return nil, fmt.Errorf("failed to batch call mint proposals: %w", err)
	}

	mintProposals := make([]pinata_bridge_abi.MintProposal, 0, mintProposalLength.Int64())
	for _, result := range results {
		if result == nil {
			continue
		}

		unpacked, err := agentCollectionABI.Unpack("mintProposal", result)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack mintProposal result: %w", err)
		}
		mintProposal := *abi.ConvertType(unpacked[0], new(pinata_bridge_abi.MintProposal)).(*pinata_bridge_abi.MintProposal)

		mintProposals = append(mintProposals, mintProposal)
	}

	return mintProposals, nil
}

func (a *AgentCollectionRequester) getAddOrUpdateAgentImageProposals(client *ethclient.Client, agentCollection *pinata_bridge_abi.AgentCollectionV1, agentCollectionAddress common.Address) ([]pinata_bridge_abi.AddOrUpdateAgentImageProposal, error) {
	addOrUpdateAgentImageProposalsLength, err := agentCollection.AddOrUpdateAgentImageProposalsLength(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get add or update image proposal length: %w", err)
	}

	agentCollectionABI, err := abi.JSON(strings.NewReader(pinata_bridge_abi.AgentCollectionV1ABI)) // You must expose raw ABI string
	if err != nil {
		return nil, fmt.Errorf("failed to parse agent collection ABI: %w", err)
	}

	batch := make([]rpc.BatchElem, 0, addOrUpdateAgentImageProposalsLength.Int64())
	results := make([]hexutil.Bytes, addOrUpdateAgentImageProposalsLength.Int64())

	for i := uint64(0); i < addOrUpdateAgentImageProposalsLength.Uint64(); i++ {
		data, err := agentCollectionABI.Pack("addOrUpdateAgentImageProposal", big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to pack addOrUpdateAgentImageProposal call: %w", err)
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

	if err := client.Client().BatchCallContext(context.Background(), batch); err != nil {
		return nil, fmt.Errorf("failed to batch call add or update agent image proposals: %w", err)
	}

	addOrUpdateAgentImageProposals := make([]pinata_bridge_abi.AddOrUpdateAgentImageProposal, 0, addOrUpdateAgentImageProposalsLength.Int64())

	for _, result := range results {
		if result == nil {
			continue
		}

		unpacked, err := agentCollectionABI.Unpack("addOrUpdateAgentImageProposal", result)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack addOrUpdateAgentImageProposal result: %w", err)
		}
		addOrUpdateAgentImageProposal := *abi.ConvertType(unpacked[0], new(pinata_bridge_abi.AddOrUpdateAgentImageProposal)).(*pinata_bridge_abi.AddOrUpdateAgentImageProposal)

		addOrUpdateAgentImageProposals = append(addOrUpdateAgentImageProposals, addOrUpdateAgentImageProposal)
	}

	return addOrUpdateAgentImageProposals, nil
}
