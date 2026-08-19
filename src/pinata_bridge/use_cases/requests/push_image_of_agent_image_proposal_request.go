package requests

type PushImageOfAgentImageProposalRequest struct {
	ProposalRequest

	AgentImageProposalId string `params:"agentImageProposalId" json:"agentImageProposalId"`
}

func (r *PushImageOfAgentImageProposalRequest) ValidateAndSanitize() error {
	if err := r.CollectionRequest.ValidateAndSanitize(); err != nil {
		return err
	}

	r.ProposalId = r.AgentImageProposalId

	return r.validateProposalId("agentImageProposalId")
}
