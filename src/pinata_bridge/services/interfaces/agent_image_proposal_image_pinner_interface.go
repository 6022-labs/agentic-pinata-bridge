package interfaces

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AgentImageProposalImagePinnerInterface pins the replacement image a pending image proposal carries.
type AgentImageProposalImagePinnerInterface interface {
	// Pin is a no-op when the proposed image is not a CID.
	Pin(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
}
