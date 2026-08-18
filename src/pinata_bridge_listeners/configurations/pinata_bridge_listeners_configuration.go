package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"go.uber.org/dig"
)

func AddPinataBridgeListenersConfiguration(container *dig.Container) {
	// Metrics
	if err := container.Provide(
		metrics.NewChainEventMetrics,
		dig.As(new(metrics_interfaces.ChainEventMetricsInterface)),
	); err != nil {
		panic(err)
	}

	// Listener
	err := container.Provide(pinata_bridge_listeners.NewAgentCollectionMintedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionMintedListener) pinata_bridge_listeners.EventListenerInterface {
		return l
	},
		dig.Group("event_listeners"),
		dig.As(new(pinata_bridge_listeners.EventListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionMintedListener) pinata_bridge_listeners.CollectionEventSubscriberInterface {
		return l
	},
		dig.Group("collection_event_subscribers"),
		dig.As(new(pinata_bridge_listeners.CollectionEventSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(pinata_bridge_listeners.NewAgentCollectionMintProposalCreatedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionMintProposalCreatedListener) pinata_bridge_listeners.EventListenerInterface {
		return l
	},
		dig.Group("event_listeners"),
		dig.As(new(pinata_bridge_listeners.EventListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionMintProposalCreatedListener) pinata_bridge_listeners.CollectionEventSubscriberInterface {
		return l
	},
		dig.Group("collection_event_subscribers"),
		dig.As(new(pinata_bridge_listeners.CollectionEventSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(pinata_bridge_listeners.NewAgentCollectionAgentImageProposalCreatedListener)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionAgentImageProposalCreatedListener) pinata_bridge_listeners.EventListenerInterface {
		return l
	},
		dig.Group("event_listeners"),
		dig.As(new(pinata_bridge_listeners.EventListenerInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(l *pinata_bridge_listeners.AgentCollectionAgentImageProposalCreatedListener) pinata_bridge_listeners.CollectionEventSubscriberInterface {
		return l
	},
		dig.Group("collection_event_subscribers"),
		dig.As(new(pinata_bridge_listeners.CollectionEventSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		pinata_bridge_listeners.NewAgentCollectionsManagerCollectionCreatedListener,
		dig.Group("event_listeners"),
		dig.As(new(pinata_bridge_listeners.EventListenerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
