package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type AgentImageProposalCreatedEventHandler struct {
	agentImageProposalImagePinner interfaces.AgentImageProposalImagePinnerInterface
}

func NewAgentImageProposalCreatedEventHandler(
	agentImageProposalImagePinner interfaces.AgentImageProposalImagePinnerInterface,
) *AgentImageProposalCreatedEventHandler {
	return &AgentImageProposalCreatedEventHandler{agentImageProposalImagePinner: agentImageProposalImagePinner}
}

func (handler *AgentImageProposalCreatedEventHandler) Handle(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1AgentImageProposalCreated,
) error {
	return handler.agentImageProposalImagePinner.Pin(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
