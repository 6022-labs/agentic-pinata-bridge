package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"go.uber.org/dig"
)

func AddPinataBridgeListenersConfiguration(container *dig.Container) {
	// Metrics
	if err := container.Provide(
		metrics.NewChainEventMetrics,
		dig.As(new(metrics_interfaces.ChainEventMetricsInterface)),
	); err != nil {
		panic(err)
	}

	// Per-collection listeners: each joins both the run group and the subscribe-on-new-collection group.
	provideCollectionListener[abi.AgentCollectionV1Minted](
		container,
		pinata_bridge_listeners.NewAgentCollectionMintedListener,
	)
	provideCollectionListener[abi.AgentCollectionV1MintProposalCreated](
		container,
		pinata_bridge_listeners.NewAgentCollectionMintProposalCreatedListener,
	)
	provideCollectionListener[abi.AgentCollectionV1AgentImageProposalCreated](
		container,
		pinata_bridge_listeners.NewAgentCollectionAgentImageProposalCreatedListener,
	)

	if err := container.Provide(
		pinata_bridge_listeners.NewAgentCollectionsManagerCollectionCreatedListener,
		dig.Group("event_listeners"),
		dig.As(new(pinata_bridge_listeners.EventListenerInterface)),
	); err != nil {
		panic(err)
	}
}
