package requests

import (
	"math/big"
	"strings"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
)

type PushMissingImagesOfAgentRequest struct {
	CollectionRequest

	AgentCollectionTokenId string `params:"agentCollectionTokenId" json:"agentCollectionTokenId"`

	agentCollectionTokenId big.Int
}

func (r *PushMissingImagesOfAgentRequest) ValidateAndSanitize() error {
	if err := r.CollectionRequest.ValidateAndSanitize(); err != nil {
		return err
	}

	if len(strings.TrimSpace(r.AgentCollectionTokenId)) == 0 {
		return errors.NewValidationError("agentCollectionTokenId", "agentCollectionTokenId is required")
	}

	tokenId, ok := big.NewInt(0).SetString(strings.TrimSpace(r.AgentCollectionTokenId), 10)
	if !ok {
		return errors.NewValidationError("agentCollectionTokenId", "agentCollectionTokenId is invalid")
	}

	r.agentCollectionTokenId = *tokenId

	return nil
}

// AgentCollectionTokenIdValue is the parsed token id; only valid after ValidateAndSanitize.
func (r *PushMissingImagesOfAgentRequest) AgentCollectionTokenIdValue() big.Int {
	return r.agentCollectionTokenId
}
