package configurations

import (
	"context"
	"net/http"

	common_metrics "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	http_pinata_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	metricnoop "go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/dig"
)

// newPinataClient builds the generated pinata client; pinata.base_url is the v3 server root.
func newPinataClient(
	pinataSettings *settings.PinataSettings,
	httpClient *http.Client,
) (clients.ClientWithResponsesInterface, error) {
	return clients.NewClientWithResponses(
		pinataSettings.BaseUrl,
		clients.WithHTTPClient(httpClient),
		clients.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+pinataSettings.ApiKey)

			return nil
		}),
	)
}

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
	err = container.Provide(newPinataClient)
	if err != nil {
		panic(err)
	}

	// Http client shared with the ipfs-check adapter; http.route separates the two upstreams.
	err = container.Provide(func() *http.Client {
		return &http.Client{
			Transport: otelhttp.NewTransport(common_metrics.NewHttpMetricsRoundTripper(
				http.DefaultTransport,
				common_metrics.NewExternalHttpMetrics("agentic_pinata_bridge_http"),
			),
				otelhttp.WithMeterProvider(metricnoop.NewMeterProvider()),
			),
		}
	})
	if err != nil {
		panic(err)
	}
}
