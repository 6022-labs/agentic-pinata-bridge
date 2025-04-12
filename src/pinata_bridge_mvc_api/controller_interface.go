package pinata_bridge_mvc_api

import "github.com/gofiber/fiber/v2"

type ControllerInterface interface {
	InitRoutes(c fiber.Router)
}
