package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_api/settings"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_event_listeners"
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

	Listeners []pinata_bridge_event_listeners.ListenerInterface `group:"listeners"`
}

// RegisterListeners
func registerListeners(p registerListenersParams) {
	for _, listener := range p.Listeners {
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

	App    *fiber.App
	Logger *zap.Logger
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
}
