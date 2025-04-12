package main

import (
	"strconv"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_api/configurations"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_api/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load(".env")

	container := configurations.ConfigureDI()
	configurations.ConfigureLogging(container)
	configurations.ConfigureServer(container)

	err := container.Invoke(func(app *fiber.App, logger *zap.Logger, httpServerSettings *settings.HttpServerSettings) error {
		logger.Info("Starting server on port " + strconv.Itoa(httpServerSettings.Port))

		return app.Listen(":" + strconv.Itoa(httpServerSettings.Port))
	})

	if err != nil {
		panic(err)
	}
}
