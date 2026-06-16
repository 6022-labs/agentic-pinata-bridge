package interfaces

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
)

type MintedEventHandlerInterface interface {
	Handle(chainId uint64, event *abi.AgentCollectionV1Minted) error
}
