package pinata_bridge_mvc

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/common/mvc"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/gofiber/fiber/v3"
)

const HealthEndpoint = "/health"

type HealthController struct {
	getHealth *use_cases.GetHealth
}

func NewHealthController(getHealth *use_cases.GetHealth) *HealthController {
	return &HealthController{getHealth: getHealth}
}

func (c *HealthController) RegisterRoutes(app fiber.Router) {
	app.Get(HealthEndpoint, c.GetHealth)
}

func (c *HealthController) GetHealth(ctx fiber.Ctx) error {
	response, err := c.getHealth.Execute(ctx.Context())
	if err != nil {
		return mvc.WriteError(ctx, err)
	}

	return ctx.JSON(response)
}
