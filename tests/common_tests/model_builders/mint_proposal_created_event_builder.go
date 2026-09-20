package model_builders

import (
	"math/big"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MintProposalCreatedEventBuilder struct {
	collectionAddress common.Address
	proposalId        *big.Int
}

func NewMintProposalCreatedEventBuilder() *MintProposalCreatedEventBuilder {
	return &MintProposalCreatedEventBuilder{
		collectionAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		proposalId:        big.NewInt(123),
	}
}

func (b *MintProposalCreatedEventBuilder) WithCollectionAddress(
	collectionAddress common.Address,
) *MintProposalCreatedEventBuilder {
	b.collectionAddress = collectionAddress
	return b
}

func (b *MintProposalCreatedEventBuilder) WithProposalId(proposalId *big.Int) *MintProposalCreatedEventBuilder {
	b.proposalId = proposalId
	return b
}

func (b *MintProposalCreatedEventBuilder) Build() *abi.AgentCollectionV1MintProposalCreated {
	return &abi.AgentCollectionV1MintProposalCreated{
		Raw:        types.Log{Address: b.collectionAddress},
		ProposalId: b.proposalId,
	}
}
