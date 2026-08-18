package requests

type PushImagesOfMintProposalRequest struct {
	ProposalRequest

	MintProposalId string `params:"mintProposalId" json:"mintProposalId"`
}

func (r *PushImagesOfMintProposalRequest) ValidateAndSanitize() error {
	if err := r.CollectionRequest.ValidateAndSanitize(); err != nil {
		return err
	}

	r.ProposalId = r.MintProposalId

	return r.validateProposalId("mintProposalId")
}
