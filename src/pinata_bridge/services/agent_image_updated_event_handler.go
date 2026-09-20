package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type AgentImageUpdatedEventHandler struct {
	agentImagesPinner interfaces.AgentImagesPinnerInterface
}

func NewAgentImageUpdatedEventHandler(
	agentImagesPinner interfaces.AgentImagesPinnerInterface,
) *AgentImageUpdatedEventHandler {
	return &AgentImageUpdatedEventHandler{agentImagesPinner: agentImagesPinner}
}

func (handler *AgentImageUpdatedEventHandler) Handle(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1AgentImageUpdated,
) error {
	return handler.agentImagesPinner.PinMissing(ctx, chainId, event.Raw.Address, *event.TokenId)
}
