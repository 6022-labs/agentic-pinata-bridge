package configurations

import (
	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics/interfaces"
	mvc_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/interfaces"
	mvc_middlewares "github.com/6022-labs/agentic-pinata-bridge/src/common/mvc/middlewares"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc"
	"go.uber.org/dig"
)

func AddPinataBridgeMvcConfiguration(container *dig.Container) {
	// Metrics
	err := container.Provide(func() metrics_interfaces.ApiRequestMetricsInterface {
		return common_metrics.NewApiRequestMetrics("agentic_pinata_bridge_api")
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(mvc_middlewares.NewApiRequestMetricsMiddleware)
	if err != nil {
		panic(err)
	}

	// Controllers
	err = container.Provide(
		pinata_bridge_mvc.NewPinataPushController,
		dig.Group("controllers"),
		dig.As(new(mvc_interfaces.ControllerInterface)),
	)
	if err != nil {
		panic(err)
	}
}
