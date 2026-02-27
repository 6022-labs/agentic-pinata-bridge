package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"go.uber.org/dig"
)

func AddPinataBridgeConfiguration(container *dig.Container) {
	// Event handlers
	err := container.Provide(
		event_handlers.NewMintedEventHandler,
		dig.As(new(event_handlers.MintedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		event_handlers.NewMintProposalCreatedEventHandler,
		dig.As(new(event_handlers.MintProposalCreatedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		event_handlers.NewAgentImageProposalCreatedEventHandler,
		dig.As(new(event_handlers.AgentImageProposalCreatedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Use cases
	err = container.Provide(
		use_cases.NewPushAgentImageCidToPinata,
		dig.As(new(use_cases.PushAgentImageCidToPinataInterface)),
	)
	if err != nil {
		panic(err)
	}
}
