package services

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type MintedEventHandler struct {
	pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface
}

func NewMintedEventHandler(pushAgentImageCidToPinata interfaces.PushAgentImageCidToPinataInterface) *MintedEventHandler {
	return &MintedEventHandler{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (h *MintedEventHandler) Handle(chainId uint64, event *abi.AgentCollectionV1Minted) error {
	return h.pushAgentImageCidToPinata.PushMissingImagesOfAgent(chainId, event.Raw.Address, *event.TokenId)
}
