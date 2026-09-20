package interfaces

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MintProposalImagesPinnerInterface pins the images a pending mint proposal carries, before approval.
type MintProposalImagesPinnerInterface interface {
	// Pin skips non-CID images; one failed pin does not stop the others.
	Pin(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error
}
