package event_handlers

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
)

type MintedEventHandlerInterface interface {
	Handle(event *abi.AgenticAIAgentCollectionMinted) error
}

type MintedEventHandler struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewMintedEventHandler(pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface) *MintedEventHandler {
	return &MintedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *MintedEventHandler) Handle(event *abi.AgenticAIAgentCollectionMinted) error {
	return h.pushAgentImageCidToPinata.PushFromAgentTokenId(*event.TokenId)
}
