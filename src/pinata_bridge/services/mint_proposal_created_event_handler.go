package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type MintProposalCreatedEventHandler struct {
	pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface
}

func NewMintProposalCreatedEventHandler(pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface) *MintProposalCreatedEventHandler {
	return &MintProposalCreatedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *MintProposalCreatedEventHandler) Handle(ctx context.Context, chainId uint64, event *abi.AgentCollectionV1MintProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImagesOfMintProposal(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
