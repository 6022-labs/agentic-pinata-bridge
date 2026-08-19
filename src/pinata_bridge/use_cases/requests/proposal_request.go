package requests

import (
	"math/big"
	"strings"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
)

// ProposalRequest addresses one proposal of one collection; the field name differs per route.
type ProposalRequest struct {
	CollectionRequest

	ProposalId string `json:"proposalId"`

	proposalId big.Int
}

func (r *ProposalRequest) ValidateAndSanitize() error {
	if err := r.CollectionRequest.ValidateAndSanitize(); err != nil {
		return err
	}

	return r.validateProposalId("proposalId")
}

// ProposalIdValue is the parsed proposal id; only valid after ValidateAndSanitize.
func (r *ProposalRequest) ProposalIdValue() big.Int {
	return r.proposalId
}

func (r *ProposalRequest) validateProposalId(field string) error {
	if len(strings.TrimSpace(r.ProposalId)) == 0 {
		return errors.NewValidationError(field, field+" is required")
	}

	proposalId, ok := big.NewInt(0).SetString(strings.TrimSpace(r.ProposalId), 10)
	if !ok {
		return errors.NewValidationError(field, field+" is invalid")
	}

	r.proposalId = *proposalId

	return nil
}
