package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_api/settings"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func ConfigureServer(container *dig.Container) {
	container.Provide(newHttpServer)

	err := container.Invoke(registerRoutes)
	if err != nil {
		panic(err)
	}
}

func newHttpServer(httpServerSettings *settings.HttpServerSettings) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: httpServerSettings.AppName,
	})
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
