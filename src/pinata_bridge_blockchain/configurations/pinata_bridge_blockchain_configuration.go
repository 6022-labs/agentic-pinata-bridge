package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/subscribers"
	blockchain_services "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/services"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
	blockchain_subscribers "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/subscribers"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
)

func AddPinataBridgeBlockchainConfiguration(container *dig.Container) {
	// Settings
	err := container.Provide(settings.NewRpcNodeSettings)
	if err != nil {
		panic(err)
	}

	err = container.Provide(settings.NewAgentCollectionsManagerSettings)
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
		blockchain_subscribers.NewAgentCollectionsManagerCollectionCreatedSubscriber,
		dig.As(new(subscribers.AgentCollectionsManagerCollectionCreatedSubscriberInterface)),
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

	// Ethereum clients
	err = container.Provide(func(rpcNodeSettings *settings.RpcNodeSettings) (*ethclient.Client, error) {
		return ethclient.Dial(rpcNodeSettings.WsUrl)
	}, dig.Name("ws"))
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(rpcNodeSettings *settings.RpcNodeSettings) (*ethclient.Client, error) {
		return ethclient.Dial(rpcNodeSettings.HttpUrl)
	}, dig.Name("http"))
	if err != nil {
		panic(err)
	}
}
