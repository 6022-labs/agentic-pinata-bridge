package interfaces

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AgentImagesPinnerInterface pins on Pinata every image of one agent that is not already there.
type AgentImagesPinnerInterface interface {
	// PinMissing skips non-CID and already-pinned images; one failed pin does not stop the others.
	PinMissing(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, tokenId big.Int) error
}
