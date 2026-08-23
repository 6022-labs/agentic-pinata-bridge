package configurations

import (
	"context"

	mvc_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/interfaces"
	mvc_middlewares "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/middlewares"
	common_settings "github.com/6022-labs/agentic-pinata-bridge/src/common/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	fiberotel "github.com/gofiber/contrib/v3/otel"
	fiberzap "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

const AppName = "6022-PinataBridge"

// ConfigureServer wires the API and the listeners; listenersContext stops the listeners on shutdown.
func ConfigureServer(container *dig.Container, listenersContext context.Context) {
	if err := container.Provide(newHttpServer); err != nil {
		panic(err)
	}

	if err := container.Provide(func() context.Context { return listenersContext }); err != nil {
		panic(err)
	}

	err := container.Invoke(useRestApi)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(useListeners)
	if err != nil {
		panic(err)
	}
}

func newHttpServer(
	hostSettings *settings.HostFeaturesSettings,
	commonHostSettings *common_settings.HostSettings,
) *fiber.App {
	if !hostSettings.UseApi {
		return nil
	}

	return fiber.New(fiber.Config{
		AppName:   AppName,
		BodyLimit: commonHostSettings.BodyLimitBytes,
	})
}

type useListenersParams struct {
	dig.In

	ListenersContext context.Context

	HostSettings   *settings.HostFeaturesSettings
	EventListeners []pinata_bridge_listeners.EventListenerInterface `group:"event_listeners"`
}

// useListeners
func useListeners(p useListenersParams) {
	if !p.HostSettings.UseListeners {
		return
	}

	ctx := p.ListenersContext

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

	HostSettings   *settings.HostFeaturesSettings
	App            *fiber.App
	Logger         *zap.Logger
	RequestMetrics *mvc_middlewares.ApiRequestMetricsMiddleware
	Controllers    []mvc_interfaces.ControllerInterface `group:"controllers"`
}

func useRestApi(p useRestApiParams) {
	if !p.HostSettings.UseApi {
		return
	}

	p.App.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// Spans only: ApiRequestMetrics owns http.server.*.
	p.App.Use(fiberotel.Middleware(fiberotel.WithoutMetrics(true)))
	p.App.Use(p.RequestMetrics.Handle)

	p.App.Use(fiberzap.New(fiberzap.Config{
		Logger: p.Logger,
	}))

	for _, controller := range p.Controllers {
		controller.RegisterRoutes(p.App)
	}
}
