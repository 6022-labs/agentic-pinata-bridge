package model_builders

import (
	"math/big"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AgentImageUpdatedEventBuilder struct {
	collectionAddress common.Address
	tokenId           *big.Int
}

func NewAgentImageUpdatedEventBuilder() *AgentImageUpdatedEventBuilder {
	return &AgentImageUpdatedEventBuilder{
		collectionAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		tokenId:           big.NewInt(123),
	}
}

func (b *AgentImageUpdatedEventBuilder) WithCollectionAddress(
	collectionAddress common.Address,
) *AgentImageUpdatedEventBuilder {
	b.collectionAddress = collectionAddress
	return b
}

func (b *AgentImageUpdatedEventBuilder) WithTokenId(tokenId *big.Int) *AgentImageUpdatedEventBuilder {
	b.tokenId = tokenId
	return b
}

func (b *AgentImageUpdatedEventBuilder) Build() *abi.AgentCollectionV1AgentImageUpdated {
	return &abi.AgentCollectionV1AgentImageUpdated{
		Raw:     types.Log{Address: b.collectionAddress},
		TokenId: b.tokenId,
	}
}
