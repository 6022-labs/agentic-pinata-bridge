package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/common/mvc"
	mvc_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/interfaces"
	mvc_middlewares "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/middlewares"
	common_settings "github.com/6022-labs/agentic-pinata-bridge/src/common/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	fiberotel "github.com/gofiber/contrib/v3/otel"
	fiberzap "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

const AppName = "6022-PinataBridge"

func ConfigureServer(container *dig.Container) {
	if err := container.Provide(newHttpServer); err != nil {
		panic(err)
	}

	err := container.Invoke(useRestApi)
	if err != nil {
		panic(err)
	}
}

func newHttpServer(
	commonHostSettings *common_settings.HostSettings,
	errorWriter *mvc.ErrorWriter,
) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:      AppName,
		BodyLimit:    commonHostSettings.BodyLimitBytes,
		ErrorHandler: errorWriter.Write,
	})
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
	p.App.Use(fiberotel.New(fiberotel.WithoutMetrics(true)))
	p.App.Use(p.RequestMetrics.Handle)

	p.App.Use(fiberzap.New(fiberzap.Config{
		Logger: p.Logger,
	}))

	for _, controller := range p.Controllers {
		controller.RegisterRoutes(p.App)
	}
}
