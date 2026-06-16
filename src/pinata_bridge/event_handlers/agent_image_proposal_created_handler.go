package event_handlers

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
)

type AgentImageProposalCreatedEventHandlerInterface interface {
	Handle(chainId uint64, event *abi.AgentCollectionV1AgentImageProposalCreated) error
}

type AgentImageProposalCreatedEventHandler struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewAgentImageProposalCreatedEventHandler(pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface) *AgentImageProposalCreatedEventHandler {
	return &AgentImageProposalCreatedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *AgentImageProposalCreatedEventHandler) Handle(chainId uint64, event *abi.AgentCollectionV1AgentImageProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImageOfAgentImageProposal(chainId, event.Raw.Address, *event.ProposalId)
}
