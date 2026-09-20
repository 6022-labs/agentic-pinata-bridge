package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
)

// PushMissingImagesOfAgent pins every image of one agent that Pinata does not already hold.
type PushMissingImagesOfAgent struct {
	agentImagesPinner interfaces.AgentImagesPinnerInterface
}

func NewPushMissingImagesOfAgent(agentImagesPinner interfaces.AgentImagesPinnerInterface) *PushMissingImagesOfAgent {
	return &PushMissingImagesOfAgent{agentImagesPinner: agentImagesPinner}
}

func (u *PushMissingImagesOfAgent) Execute(
	ctx context.Context,
	request *requests.PushMissingImagesOfAgentRequest,
) (*responses.PushResponse, error) {
	if err := request.ValidateAndSanitize(); err != nil {
		return nil, err
	}

	if err := u.agentImagesPinner.PinMissing(
		ctx,
		request.ChainIdValue(),
		request.AgentCollectionAddressValue(),
		request.AgentCollectionTokenIdValue(),
	); err != nil {
		return nil, err
	}

	return &responses.PushResponse{}, nil
}
