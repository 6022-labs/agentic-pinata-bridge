package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"
	ipfs_check_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/settings"
	"go.uber.org/dig"
)

func AddPinataBridgeHttpIpfsCheckConfiguration(container *dig.Container) {
	// Settings
	err := container.Provide(settings.NewIpfsCheckSettings)
	if err != nil {
		panic(err)
	}

	// Clients
	err = container.Provide(
		clients.NewIpfsCheckClient,
		dig.As(new(clients.IpfsCheckClientInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Services
	err = container.Provide(
		ipfs_check_services.NewIpfsCheckRequester,
		dig.As(new(services.IpfsCheckRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}
}
