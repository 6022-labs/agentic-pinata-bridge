package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
)

// HandleAgentImageProposalCreatedEvent pins the image carried by a freshly created agent-image proposal.
type HandleAgentImageProposalCreatedEvent struct {
	pushImageOfAgentImageProposal *PushImageOfAgentImageProposal
}

func NewHandleAgentImageProposalCreatedEvent(
	pushImageOfAgentImageProposal *PushImageOfAgentImageProposal,
) *HandleAgentImageProposalCreatedEvent {
	return &HandleAgentImageProposalCreatedEvent{pushImageOfAgentImageProposal: pushImageOfAgentImageProposal}
}

func (u *HandleAgentImageProposalCreatedEvent) Execute(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1AgentImageProposalCreated,
) error {
	return u.pushImageOfAgentImageProposal.push(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
