package configurations

import (
	"net/http"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	http_pinata_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
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

	// Http clients
	err = container.Provide(func() *http.Client {
		return http.DefaultClient
	})
	if err != nil {
		panic(err)
	}
}
