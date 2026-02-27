package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_event_listeners"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func ConfigureServer(container *dig.Container) {
	container.Provide(newHttpServer)

	err := container.Invoke(useRestApi)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(useListeners)
	if err != nil {
		panic(err)
	}
}

func newHttpServer(hostSettings *settings.HostSettings) *fiber.App {
	if !hostSettings.UseApi {
		return nil
	}

	return fiber.New(fiber.Config{
		AppName: hostSettings.AppName,
	})
}

type useListenersParams struct {
	dig.In

	HostSettings                     *settings.HostSettings
	AgentCollectionsManagerRequester services.AgentCollectionsManagerRequesterInterface
	Listeners                        []pinata_bridge_event_listeners.ListenerInterface `group:"listeners"`
}

// RegisterListeners
func useListeners(p useListenersParams) {
	if !p.HostSettings.UseListeners {
		return
	}

	collectionAddresses, err := p.AgentCollectionsManagerRequester.GetAllCollectionAddresses()
	if err != nil {
		panic(err)
	}

	for _, listener := range p.Listeners {
		// If the listener is a CollectionListenerInterface, subscribe to the collections
		if collectionListener, ok := listener.(pinata_bridge_event_listeners.CollectionListenerInterface); ok {
			for _, collection := range collectionAddresses {
				err := collectionListener.Subscribe(collection)
				if err != nil {
					panic(err)
				}
			}
		}

		go func() {
			err := listener.Listen()
			if err != nil {
				panic(err)
			}
		}()
	}
}

type useRestApiParams struct {
	dig.In

	HostSettings *settings.HostSettings
	App          *fiber.App
	Logger       *zap.Logger
	Controllers  []pinata_bridge_mvc.ControllerInterface `group:"controllers"`
}

// UseRestApi hooks up the routes and uses Fx to create new controller instances per request
func useRestApi(p useRestApiParams) {
	if !p.HostSettings.UseApi {
		return
	}

	p.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	p.App.Use(fiberzap.New(fiberzap.Config{
		Logger: p.Logger,
	}))

	for _, controller := range p.Controllers {
		controller.InitRoutes(p.App)
	}
}
