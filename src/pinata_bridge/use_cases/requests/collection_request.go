package requests

import (
	"strconv"
	"strings"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/ethereum/go-ethereum/common"
)

// CollectionRequest is the chain + collection pair every push request carries.
type CollectionRequest struct {
	ChainId                string `params:"chainId"                json:"chainId"`
	AgentCollectionAddress string `params:"agentCollectionAddress" json:"agentCollectionAddress"`

	chainId                uint64
	agentCollectionAddress common.Address
}

func (r *CollectionRequest) ValidateAndSanitize() error {
	if len(strings.TrimSpace(r.ChainId)) == 0 {
		return errors.NewValidationError("chainId", "chainId is required")
	}

	chainId, err := strconv.ParseUint(strings.TrimSpace(r.ChainId), 10, 64)
	if err != nil {
		return errors.NewValidationError("chainId", "chainId is invalid")
	}

	address := strings.TrimSpace(r.AgentCollectionAddress)
	if len(address) == 0 {
		return errors.NewValidationError("agentCollectionAddress", "agentCollectionAddress is required")
	}

	if !common.IsHexAddress(address) {
		return errors.NewValidationError("agentCollectionAddress", "agentCollectionAddress is invalid")
	}

	r.chainId = chainId
	r.agentCollectionAddress = common.HexToAddress(address)

	return nil
}

// ChainIdValue is the parsed chain id; only valid after ValidateAndSanitize.
func (r *CollectionRequest) ChainIdValue() uint64 {
	return r.chainId
}

// AgentCollectionAddressValue is the parsed collection address; only valid after ValidateAndSanitize.
func (r *CollectionRequest) AgentCollectionAddressValue() common.Address {
	return r.agentCollectionAddress
}
