package event_handlers

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
)

type MintedEventHandlerInterface interface {
	Handle(chainId uint64, event *abi.AgentCollectionV1Minted) error
}

type MintedEventHandler struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewMintedEventHandler(pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface) *MintedEventHandler {
	return &MintedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *MintedEventHandler) Handle(chainId uint64, event *abi.AgentCollectionV1Minted) error {
	return h.pushAgentImageCidToPinata.PushMissingImagesOfAgent(chainId, event.Raw.Address, *event.TokenId)
}
