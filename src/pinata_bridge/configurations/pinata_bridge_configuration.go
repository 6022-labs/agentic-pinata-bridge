package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"go.uber.org/dig"
)

func AddPinataBridgeConfiguration(container *dig.Container) {
	// Settings
	err := container.Provide(settings.NewChainsSettings)
	if err != nil {
		panic(err)
	}

	// Event handlers
	err = container.Provide(
		services.NewMintedEventHandler,
		dig.As(new(interfaces.MintedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		services.NewMintProposalCreatedEventHandler,
		dig.As(new(interfaces.MintProposalCreatedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		services.NewAgentImageProposalCreatedEventHandler,
		dig.As(new(interfaces.AgentImageProposalCreatedEventHandlerInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Use cases
	err = container.Provide(
		services.NewPushAgentImageCidToPinata,
		dig.As(new(interfaces.PushAgentImageCidToPinataInterface)),
	)
	if err != nil {
		panic(err)
	}
}
