package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	"go.uber.org/dig"
)

// provideCollectionListener registers one listener three times: concrete, as a run-loop
// listener, and as a subscriber the collection-created listener notifies about new collections.
func provideCollectionListener[T pinata_bridge_listeners.ChainEvent](container *dig.Container, constructor any) {
	if err := container.Provide(constructor); err != nil {
		panic(err)
	}

	if err := container.Provide(
		func(listener *pinata_bridge_listeners.ChainEventListener[T]) pinata_bridge_listeners.EventListenerInterface {
			return listener
		},
		dig.Group("event_listeners"),
	); err != nil {
		panic(err)
	}

	if err := container.Provide(
		func(listener *pinata_bridge_listeners.ChainEventListener[T]) pinata_bridge_listeners.CollectionEventSubscriberInterface {
			return listener
		},
		dig.Group("collection_event_subscribers"),
	); err != nil {
		panic(err)
	}
}
