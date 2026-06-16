package event_handlers

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
)

type MintProposalCreatedEventHandlerInterface interface {
	Handle(chainId uint64, event *abi.AgentCollectionV1MintProposalCreated) error
}

type MintProposalCreatedEventHandler struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewMintProposalCreatedEventHandler(pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface) *MintProposalCreatedEventHandler {
	return &MintProposalCreatedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *MintProposalCreatedEventHandler) Handle(chainId uint64, event *abi.AgentCollectionV1MintProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImagesOfMintProposal(chainId, event.Raw.Address, *event.ProposalId)
}
