package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/subscribers"
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
		dig.As(new(subscribers.AgentCollectionMintedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionMintProposalCreatedSubscriber,
		dig.As(new(subscribers.AgentCollectionMintProposalCreatedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionsManagerCollectionCreatedSubscriber,
		dig.As(new(subscribers.AgentCollectionsManagerCollectionCreatedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_subscribers.NewAgentCollectionAgentImageProposalCreatedSubscriber,
		dig.As(new(subscribers.AgentCollectionAgentImageProposalCreatedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Services
	err = container.Provide(
		blockchain_services.NewAgentCollectionRequester,
		dig.As(new(services.AgentCollectionRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		blockchain_services.NewAgentCollectionsManagerRequester,
		dig.As(new(services.AgentCollectionsManagerRequesterInterface)),
	)
	if err != nil {
		panic(err)
	}

}
