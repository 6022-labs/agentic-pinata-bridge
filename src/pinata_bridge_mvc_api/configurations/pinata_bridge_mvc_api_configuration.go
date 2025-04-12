package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_mvc_api"
	"go.uber.org/dig"
)

func AddPinataBridgeMvcApiConfiguration(container *dig.Container) {
	// Controllers
	err := container.Provide(
		pinata_bridge_mvc_api.NewPinataPushController,
		dig.Group("controllers"),
		dig.As(new(pinata_bridge_mvc_api.ControllerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
