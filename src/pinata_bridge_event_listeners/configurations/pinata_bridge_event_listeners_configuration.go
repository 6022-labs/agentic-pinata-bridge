package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_event_listeners"
	"go.uber.org/dig"
)

func AddPinataBridgeEventListenersConfiguration(container *dig.Container) {
	// Listener
	err := container.Provide(pinata_bridge_event_listeners.NewAgentCollectionMintedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionMintedListener) pinata_bridge_event_listeners.ListenerInterface {
		return l
	},
		dig.Group("listeners"),
		dig.As(new(pinata_bridge_event_listeners.ListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionMintedListener) pinata_bridge_event_listeners.CollectionListenerInterface {
		return l
	},
		dig.Group("collection_listeners"),
		dig.As(new(pinata_bridge_event_listeners.CollectionListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(pinata_bridge_event_listeners.NewAgentCollectionMintProposalCreatedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionMintProposalCreatedListener) pinata_bridge_event_listeners.ListenerInterface {
		return l
	},
		dig.Group("listeners"),
		dig.As(new(pinata_bridge_event_listeners.ListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionMintProposalCreatedListener) pinata_bridge_event_listeners.CollectionListenerInterface {
		return l
	},
		dig.Group("collection_listeners"),
		dig.As(new(pinata_bridge_event_listeners.CollectionListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(pinata_bridge_event_listeners.NewAgentCollectionAgentImageProposalCreatedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionAgentImageProposalCreatedListener) pinata_bridge_event_listeners.ListenerInterface {
		return l
	},
		dig.Group("listeners"),
		dig.As(new(pinata_bridge_event_listeners.ListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_event_listeners.AgentCollectionAgentImageProposalCreatedListener) pinata_bridge_event_listeners.CollectionListenerInterface {
		return l
	},
		dig.Group("collection_listeners"),
		dig.As(new(pinata_bridge_event_listeners.CollectionListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		pinata_bridge_event_listeners.NewAgentCollectionsManagerCollectionCreatedListener,
		dig.Group("listeners"),
		dig.As(new(pinata_bridge_event_listeners.ListenerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
