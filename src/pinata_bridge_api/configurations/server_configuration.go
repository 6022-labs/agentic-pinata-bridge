package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_api/settings"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_event_listeners"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_mvc_api"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func ConfigureServer(container *dig.Container) {
	container.Provide(newHttpServer)

	err := container.Invoke(registerListeners)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(registerRoutes)
	if err != nil {
		panic(err)
	}
}

func newHttpServer(httpServerSettings *settings.HttpServerSettings) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: httpServerSettings.AppName,
	})
}

type registerListenersParams struct {
	dig.In

	AgentCollectionsManagerRequester services.AgentCollectionsManagerRequesterInterface
	Listeners                        []pinata_bridge_event_listeners.ListenerInterface `group:"listeners"`
}

// RegisterListeners
func registerListeners(p registerListenersParams) {
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

type registerRoutesParams struct {
	dig.In

	App         *fiber.App
	Logger      *zap.Logger
	Controllers []pinata_bridge_mvc_api.ControllerInterface `group:"controllers"`
}

// RegisterRoutes hooks up the routes and uses Fx to create new controller instances per request
func registerRoutes(p registerRoutesParams) {
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
