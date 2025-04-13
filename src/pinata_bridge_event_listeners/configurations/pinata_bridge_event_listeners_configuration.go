package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_event_listeners"
	"go.uber.org/dig"
)

func AddPinataBridgeEventListenersConfiguration(container *dig.Container) {
	// Listeners
	err := container.Provide(
		pinata_bridge_event_listeners.NewAgenticAIAgentCollectionMintedListener,
		dig.Group("listeners"),
		dig.As(new(pinata_bridge_event_listeners.ListenerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
