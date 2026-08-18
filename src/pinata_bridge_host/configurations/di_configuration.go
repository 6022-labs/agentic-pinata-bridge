package configurations

import (
	common_settings "github.com/6022-labs/agentic-pinata-bridge/src/common/settings"
	pinata_bridge_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/configurations"
	pinata_bridge_blockchain_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/configurations"
	pinata_bridge_host_settings "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	pinata_bridge_http_ipfs_check_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/configurations"
	pinata_bridge_http_pinata_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/configurations"
	pinata_bridge_listeners_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/configurations"
	pinata_bridge_mvc_configuration "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc/configurations"
	"github.com/knadh/koanf/v2"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

const (
	hostSettingsKey = "host"
	defaultApiPort  = 3000
)

func ConfigureDI(config *koanf.Koanf) *dig.Container {
	container := dig.New()

	if err := container.Provide(func() *koanf.Koanf {
		return config
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(pinata_bridge_host_settings.NewHostFeaturesSettings); err != nil {
		panic(err)
	}

	if err := container.Provide(func(logger *zap.Logger, k *koanf.Koanf) *common_settings.HostSettings {
		return common_settings.NewHostSettings(logger, k, hostSettingsKey, defaultApiPort)
	}); err != nil {
		panic(err)
	}

	pinata_bridge_configuration.AddPinataBridgeConfiguration(container)
	pinata_bridge_mvc_configuration.AddPinataBridgeMvcConfiguration(container)
	pinata_bridge_blockchain_configuration.AddPinataBridgeBlockchainConfiguration(container)
	pinata_bridge_http_pinata_configuration.AddPinataBridgeHttpPinataConfiguration(container)
	pinata_bridge_http_ipfs_check_configuration.AddPinataBridgeHttpIpfsCheckConfiguration(container)
	pinata_bridge_listeners_configuration.AddPinataBridgeListenersConfiguration(container)

	return container
}
