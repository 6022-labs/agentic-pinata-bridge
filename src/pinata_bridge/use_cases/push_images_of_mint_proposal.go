package use_cases

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/responses"
)

// PushImagesOfMintProposal pins the images of a pending mint proposal so they resolve before approval.
type PushImagesOfMintProposal struct {
	mintProposalImagesPinner interfaces.MintProposalImagesPinnerInterface
}

func NewPushImagesOfMintProposal(
	mintProposalImagesPinner interfaces.MintProposalImagesPinnerInterface,
) *PushImagesOfMintProposal {
	return &PushImagesOfMintProposal{mintProposalImagesPinner: mintProposalImagesPinner}
}

func (u *PushImagesOfMintProposal) Execute(
	ctx context.Context,
	request *requests.PushImagesOfMintProposalRequest,
) (*responses.PushResponse, error) {
	if err := request.ValidateAndSanitize(); err != nil {
		return nil, err
	}

	if err := u.mintProposalImagesPinner.Pin(
		ctx,
		request.ChainIdValue(),
		request.AgentCollectionAddressValue(),
		request.ProposalIdValue(),
	); err != nil {
		return nil, err
	}

	return &responses.PushResponse{}, nil
}
