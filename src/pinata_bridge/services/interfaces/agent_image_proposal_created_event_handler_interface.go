package interfaces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
)

type AgentImageProposalCreatedEventHandlerInterface interface {
	Handle(ctx context.Context, chainId uint64, event *abi.AgentCollectionV1AgentImageProposalCreated) error
}
