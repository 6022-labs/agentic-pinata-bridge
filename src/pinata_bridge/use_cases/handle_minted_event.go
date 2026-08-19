package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
)

// HandleMintedEvent pins the images of an agent as soon as it is minted.
type HandleMintedEvent struct {
	pushMissingImagesOfAgent *PushMissingImagesOfAgent
}

func NewHandleMintedEvent(pushMissingImagesOfAgent *PushMissingImagesOfAgent) *HandleMintedEvent {
	return &HandleMintedEvent{pushMissingImagesOfAgent: pushMissingImagesOfAgent}
}

func (u *HandleMintedEvent) Execute(ctx context.Context, chainId uint64, event *abi.AgentCollectionV1Minted) error {
	return u.pushMissingImagesOfAgent.push(ctx, chainId, event.Raw.Address, *event.TokenId)
}
