package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
)

// HandleAgentImageUpdatedEvent re-sweeps an agent's images after any on-chain image write.
type HandleAgentImageUpdatedEvent struct {
	pushMissingImagesOfAgent *PushMissingImagesOfAgent
}

func NewHandleAgentImageUpdatedEvent(pushMissingImagesOfAgent *PushMissingImagesOfAgent) *HandleAgentImageUpdatedEvent {
	return &HandleAgentImageUpdatedEvent{pushMissingImagesOfAgent: pushMissingImagesOfAgent}
}

func (u *HandleAgentImageUpdatedEvent) Execute(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1AgentImageUpdated,
) error {
	return u.pushMissingImagesOfAgent.push(ctx, chainId, event.Raw.Address, *event.TokenId)
}
