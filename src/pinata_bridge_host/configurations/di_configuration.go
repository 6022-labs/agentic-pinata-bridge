package configurations

import (
	pinata_bridge_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/configurations"
	pinata_bridge_blockchain_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/configurations"
	pinata_bridge_event_listeners_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_event_listeners/configurations"
	pinata_bridge_host_settings "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	pinata_bridge_http_ipfs_check_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/configurations"
	pinata_bridge_http_pinata_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/configurations"
	pinata_bridge_mvc_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc/configurations"
	"go.uber.org/dig"
)

func ConfigureDI() *dig.Container {
	container := dig.New()

	container.Provide(pinata_bridge_host_settings.NewHostSettings)

	pinata_bridge_configuration.AddPinataBridgeConfiguration(container)
	pinata_bridge_mvc_configuration.AddPinataBridgeMvcConfiguration(container)
	pinata_bridge_blockchain_configuration.AddPinataBridgeBlockchainConfiguration(container)
	pinata_bridge_http_pinata_configuration.AddPinataBridgeHttpPinataConfiguration(container)
	pinata_bridge_http_ipfs_check_configuration.AddPinataBridgeHttpIpfsCheckConfiguration(container)
	pinata_bridge_event_listeners_configuration.AddPinataBridgeEventListenersConfiguration(container)

	return container
}
