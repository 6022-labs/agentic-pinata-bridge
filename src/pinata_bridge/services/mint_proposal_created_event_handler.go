package services

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
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

func (h *MintProposalCreatedEventHandler) Handle(chainId uint64, event *abi.AgentCollectionV1MintProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImagesOfMintProposal(chainId, event.Raw.Address, *event.ProposalId)
}
