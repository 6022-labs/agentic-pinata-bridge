package model_builders

import (
	"math/big"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MintedEventBuilder struct {
	collectionAddress common.Address
	tokenId           *big.Int
}

func NewMintedEventBuilder() *MintedEventBuilder {
	return &MintedEventBuilder{
		collectionAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		tokenId:           big.NewInt(123),
	}
}

func (b *MintedEventBuilder) WithCollectionAddress(collectionAddress common.Address) *MintedEventBuilder {
	b.collectionAddress = collectionAddress
	return b
}

func (b *MintedEventBuilder) WithTokenId(tokenId *big.Int) *MintedEventBuilder {
	b.tokenId = tokenId
	return b
}

func (b *MintedEventBuilder) Build() *abi.AgentCollectionV1Minted {
	return &abi.AgentCollectionV1Minted{
		Raw:     types.Log{Address: b.collectionAddress},
		TokenId: b.tokenId,
	}
}
