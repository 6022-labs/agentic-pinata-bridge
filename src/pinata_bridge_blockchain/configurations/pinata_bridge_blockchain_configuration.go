package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/metrics/interfaces"
	blockchain_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"go.uber.org/dig"
)

func AddPinataBridgeBlockchainConfiguration(container *dig.Container) {
	// Metrics
	if err := container.Provide(
		metrics.NewBlockchainRpcMetrics,
		dig.As(new(metrics_interfaces.BlockchainRpcMetricsInterface)),
	); err != nil {
		panic(err)
	}

	// Settings
	err := container.Provide(settings.NewRpcSettings)
	if err != nil {
		panic(err)
	}

	err = container.Provide(settings.NewAgentCollectionsManagersSettings)
	if err != nil {
		panic(err)
	}

	// Eth client factory
	err = container.Provide(factory.NewEthClientFactory)
	if err != nil {
		panic(err)
	}

	// Event subscription providers
	err = container.Provide(
		blockchain_services.NewAgentCollectionMintedEventSubscriptionProvider,
		dig.As(new(interfaces.AgentCollectionMintedEventSubscriptionProviderInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_services.NewAgentCollectionMintProposalCreatedEventSubscriptionProvider,
		dig.As(new(interfaces.AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_services.NewAgentCollectionsManagerCollectionCreatedEventSubscriptionProvider,
		dig.As(new(interfaces.AgentCollectionsManagerCollectionCreatedEventSubscriptionProviderInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_services.NewAgentCollectionAgentImageProposalCreatedEventSubscriptionProvider,
		dig.As(new(interfaces.AgentCollectionAgentImageProposalCreatedEventSubscriptionProviderInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Services
	err = container.Provide(
		blockchain_services.NewAgentCollectionRequester,
		dig.As(new(interfaces.AgentCollectionRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_services.NewAgentCollectionsManagerRequester,
		dig.As(new(interfaces.AgentCollectionsManagerRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}

}
