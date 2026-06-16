package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	blockchain_services "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	blockchain_subscribers "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/subscribers"
	"go.uber.org/dig"
)

func AddPinataBridgeBlockchainConfiguration(container *dig.Container) {
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

	// Event subscribers
	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionMintedSubscriber,
		dig.As(new(interfaces.AgentCollectionMintedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionMintProposalCreatedSubscriber,
		dig.As(new(interfaces.AgentCollectionMintProposalCreatedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionsManagerCollectionCreatedSubscriber,
		dig.As(new(interfaces.AgentCollectionsManagerCollectionCreatedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionAgentImageProposalCreatedSubscriber,
		dig.As(new(interfaces.AgentCollectionAgentImageProposalCreatedSubscriberInterface)),
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
