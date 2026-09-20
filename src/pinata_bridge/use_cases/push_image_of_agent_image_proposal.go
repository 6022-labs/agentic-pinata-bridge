package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
)

// PushImageOfAgentImageProposal pins the replacement image of a pending agent-image proposal.
type PushImageOfAgentImageProposal struct {
	agentImageProposalImagePinner interfaces.AgentImageProposalImagePinnerInterface
}

func NewPushImageOfAgentImageProposal(
	agentImageProposalImagePinner interfaces.AgentImageProposalImagePinnerInterface,
) *PushImageOfAgentImageProposal {
	return &PushImageOfAgentImageProposal{agentImageProposalImagePinner: agentImageProposalImagePinner}
}

func (u *PushImageOfAgentImageProposal) Execute(
	ctx context.Context,
	request *requests.PushImageOfAgentImageProposalRequest,
) (*responses.PushResponse, error) {
	if err := request.ValidateAndSanitize(); err != nil {
		return nil, err
	}

	if err := u.agentImageProposalImagePinner.Pin(
		ctx,
		request.ChainIdValue(),
		request.AgentCollectionAddressValue(),
		request.ProposalIdValue(),
	); err != nil {
		return nil, err
	}

	return &responses.PushResponse{}, nil
}
