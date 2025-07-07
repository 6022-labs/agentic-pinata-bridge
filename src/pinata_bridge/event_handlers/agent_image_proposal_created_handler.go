package event_handlers

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
)

type AgentImageProposalCreatedEventHandlerInterface interface {
	Handle(event *abi.AgentCollectionV1AgentImageProposalCreated) error
}

type AgentImageProposalCreatedEventHandler struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewAgentImageProposalCreatedEventHandler(pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface) *AgentImageProposalCreatedEventHandler {
	return &AgentImageProposalCreatedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *AgentImageProposalCreatedEventHandler) Handle(event *abi.AgentCollectionV1AgentImageProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImageOfAgentImageProposal(event.Raw.Address, *event.ProposalId)
}
