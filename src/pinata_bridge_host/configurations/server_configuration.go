package configurations

import (
	"context"

	mvc_middlewares "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/middlewares"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	mvc_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/interfaces"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

const AppName = "6022-PinataBridge"

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
		AppName: AppName,
	})
}

type useListenersParams struct {
	dig.In

	HostSettings   *settings.HostSettings
	EventListeners []pinata_bridge_listeners.EventListenerInterface `group:"event_listeners"`
}

// RegisterListeners
func useListeners(p useListenersParams) {
	if !p.HostSettings.UseListeners {
		return
	}

	ctx := context.Background()

	for _, listener := range p.EventListeners {
		if err := listener.SubscribeAll(ctx); err != nil {
			panic(err)
		}

		go func(listener pinata_bridge_listeners.EventListenerInterface) {
			if err := listener.Listen(ctx); err != nil {
				panic(err)
			}
		}(listener)
	}
}

type useRestApiParams struct {
	dig.In

	HostSettings   *settings.HostSettings
	App            *fiber.App
	Logger         *zap.Logger
	RequestMetrics *mvc_middlewares.ApiRequestMetricsMiddleware
	Controllers    []mvc_interfaces.ControllerInterface `group:"controllers"`
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

	// Spans only: ApiRequestMetrics owns http.server.*.
	p.App.Use(otelfiber.Middleware(otelfiber.WithoutMetrics(true)))
	p.App.Use(p.RequestMetrics.Handle)

	p.App.Use(fiberzap.New(fiberzap.Config{
		Logger: p.Logger,
	}))

	for _, controller := range p.Controllers {
		controller.RegisterRoutes(p.App)
	}
}
