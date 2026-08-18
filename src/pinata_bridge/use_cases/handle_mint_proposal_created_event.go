package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
)

// HandleMintProposalCreatedEvent pins the images carried by a freshly created mint proposal.
type HandleMintProposalCreatedEvent struct {
	pushImagesOfMintProposal *PushImagesOfMintProposal
}

func NewHandleMintProposalCreatedEvent(
	pushImagesOfMintProposal *PushImagesOfMintProposal,
) *HandleMintProposalCreatedEvent {
	return &HandleMintProposalCreatedEvent{pushImagesOfMintProposal: pushImagesOfMintProposal}
}

func (u *HandleMintProposalCreatedEvent) Execute(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1MintProposalCreated,
) error {
	return u.pushImagesOfMintProposal.push(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
