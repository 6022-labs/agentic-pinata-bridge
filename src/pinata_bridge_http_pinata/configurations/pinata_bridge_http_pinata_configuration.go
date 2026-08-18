package configurations

import (
	"net/http"

	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	http_pinata_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/dig"
)

func AddPinataBridgeHttpPinataConfiguration(container *dig.Container) {
	// Settings
	err := container.Provide(settings.NewPinataSettings)
	if err != nil {
		panic(err)
	}

	// Services
	err = container.Provide(
		http_pinata_services.NewPinataRequester,
		dig.As(new(interfaces.PinataRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Clients
	err = container.Provide(
		clients.NewPinataClient,
		dig.As(new(clients.PinataClientInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Http client shared with the ipfs-check adapter; http.route separates the two upstreams.
	err = container.Provide(func() *http.Client {
		return &http.Client{
			Transport: otelhttp.NewTransport(common_metrics.NewHttpMetricsRoundTripper(
				http.DefaultTransport,
				common_metrics.NewExternalHttpMetrics("agentic_pinata_bridge_http"),
			)),
		}
	})
	if err != nil {
		panic(err)
	}
}
