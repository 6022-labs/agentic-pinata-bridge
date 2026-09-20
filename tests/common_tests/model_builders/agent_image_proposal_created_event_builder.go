package model_builders

import (
	"math/big"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AgentImageProposalCreatedEventBuilder struct {
	collectionAddress common.Address
	proposalId        *big.Int
}

func NewAgentImageProposalCreatedEventBuilder() *AgentImageProposalCreatedEventBuilder {
	return &AgentImageProposalCreatedEventBuilder{
		collectionAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		proposalId:        big.NewInt(123),
	}
}

func (b *AgentImageProposalCreatedEventBuilder) WithCollectionAddress(
	collectionAddress common.Address,
) *AgentImageProposalCreatedEventBuilder {
	b.collectionAddress = collectionAddress
	return b
}

func (b *AgentImageProposalCreatedEventBuilder) WithProposalId(
	proposalId *big.Int,
) *AgentImageProposalCreatedEventBuilder {
	b.proposalId = proposalId
	return b
}

func (b *AgentImageProposalCreatedEventBuilder) Build() *abi.AgentCollectionV1AgentImageProposalCreated {
	return &abi.AgentCollectionV1AgentImageProposalCreated{
		Raw:        types.Log{Address: b.collectionAddress},
		ProposalId: b.proposalId,
	}
}
