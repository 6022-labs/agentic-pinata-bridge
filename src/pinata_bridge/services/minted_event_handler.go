package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type MintedEventHandler struct {
	agentImagesPinner interfaces.AgentImagesPinnerInterface
}

func NewMintedEventHandler(agentImagesPinner interfaces.AgentImagesPinnerInterface) *MintedEventHandler {
	return &MintedEventHandler{agentImagesPinner: agentImagesPinner}
}

func (handler *MintedEventHandler) Handle(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1Minted,
) error {
	return handler.agentImagesPinner.PinMissing(ctx, chainId, event.Raw.Address, *event.TokenId)
}
