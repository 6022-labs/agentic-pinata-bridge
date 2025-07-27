package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_mvc"
	"go.uber.org/dig"
)

func AddPinataBridgeMvcConfiguration(container *dig.Container) {
	// Controllers
	err := container.Provide(
		pinata_bridge_mvc.NewPinataPushController,
		dig.Group("controllers"),
		dig.As(new(pinata_bridge_mvc.ControllerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
