package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
)

type MintProposalCreatedEventHandler struct {
	mintProposalImagesPinner interfaces.MintProposalImagesPinnerInterface
}

func NewMintProposalCreatedEventHandler(
	mintProposalImagesPinner interfaces.MintProposalImagesPinnerInterface,
) *MintProposalCreatedEventHandler {
	return &MintProposalCreatedEventHandler{mintProposalImagesPinner: mintProposalImagesPinner}
}

func (handler *MintProposalCreatedEventHandler) Handle(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionV1MintProposalCreated,
) error {
	return handler.mintProposalImagesPinner.Pin(ctx, chainId, event.Raw.Address, *event.ProposalId)
}
