package configurations

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/event_subscribers"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/services"
	blockchain_event_subscribers "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/event_subscribers"
	blockchain_services "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/services"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
)

func AddPinataBridgeBlockchainConfiguration(container *dig.Container) {
	// Settings
	err := container.Provide(settings.NewRpcNodeSettings)
	if err != nil {
		panic(err)
	}

	err = container.Provide(settings.NewAgenticAIAgentCollectionSettings)
	if err != nil {
		panic(err)
	}

	// Event subscribers
	err = container.Provide(
		blockchain_event_subscribers.NewAgenticAIAgentCollectionMintedSubscriber,
		dig.As(new(event_subscribers.AgenticAIAgentCollectionMintedSubscriberInterface)),
	)
	if err != nil {
		panic(err)
	}

	// Services
	err = container.Provide(
		blockchain_services.NewAgenticAIAgentCollectionRequester,
		dig.As(new(services.AgenticAIAgentCollectionRequesterInterface)),
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
