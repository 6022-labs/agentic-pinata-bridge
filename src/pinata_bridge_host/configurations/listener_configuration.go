package configurations

import (
	"context"
	"sync"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/dig"
)

type configureListenersParams struct {
	dig.In

	App            *fiber.App
	HostSettings   *settings.HostFeaturesSettings
	EventListeners []pinata_bridge_listeners.EventListenerInterface `group:"event_listeners"`
}

// ConfigureListeners starts the chain listeners and stops them before Fiber shuts down.
func ConfigureListeners(container *dig.Container) {
	err := container.Invoke(func(p configureListenersParams) error {
		if !p.HostSettings.UseListeners {
			return nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		listenersWg := &sync.WaitGroup{}
		p.App.Hooks().OnPreShutdown(func() error {
			cancel()
			listenersWg.Wait()
			return nil
		})

		for _, listener := range p.EventListeners {
			if err := listener.SubscribeAll(ctx); err != nil {
				return err
			}

			listenersWg.Go(func() {
				// A subscription may fail while it is being torn down; only a live one is fatal.
				if err := listener.Listen(ctx); err != nil && ctx.Err() == nil {
					panic(err)
				}
			})
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
