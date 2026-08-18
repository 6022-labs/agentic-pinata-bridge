package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type AgentImageProposalCreatedEventHandler struct {
	pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface
}

func NewAgentImageProposalCreatedEventHandler(pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface) *AgentImageProposalCreatedEventHandler {
	return &AgentImageProposalCreatedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *AgentImageProposalCreatedEventHandler) Handle(ctx context.Context, chainId uint64, event *abi.AgentCollectionV1AgentImageProposalCreated) error {
	return h.pushAgentImageCidToPinata.PushImageOfAgentImageProposal(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
