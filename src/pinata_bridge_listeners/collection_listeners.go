package pinata_bridge_listeners

import (

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"go.uber.org/zap"
)

type (
	AgentCollectionMintedListener              = ChainEventListener[abi.AgentCollectionV1Minted]
	AgentCollectionMintProposalCreatedListener = ChainEventListener[abi.AgentCollectionV1MintProposalCreated]
	// AgentCollectionAgentImageProposalCreatedListener watches proposals that replace an agent's image.
	AgentCollectionAgentImageProposalCreatedListener = ChainEventListener[abi.AgentCollectionV1AgentImageProposalCreated]
)

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	subscriptionProvider interfaces.MintedSubscriptionProviderInterface,
	handleMintedEvent *use_cases.HandleMintedEvent,
) *AgentCollectionMintedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.Minted",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		subscriptionProvider.StartMintedSubscription,
		handleMintedEvent.Execute,
	)
}

func NewAgentCollectionMintProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	subscriptionProvider interfaces.MintProposalCreatedSubscriptionProviderInterface,
	handleMintProposalCreatedEvent *use_cases.HandleMintProposalCreatedEvent,
) *AgentCollectionMintProposalCreatedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.MintProposalCreated",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		subscriptionProvider.StartMintProposalCreatedSubscription,
		handleMintProposalCreatedEvent.Execute,
	)
}

func NewAgentCollectionAgentImageProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	subscriptionProvider interfaces.AgentImageProposalCreatedSubscriptionProviderInterface,
	handleAgentImageProposalCreatedEvent *use_cases.HandleAgentImageProposalCreatedEvent,
) *AgentCollectionAgentImageProposalCreatedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.AgentImageProposalCreated",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		subscriptionProvider.StartAgentImageProposalCreatedSubscription,
		handleAgentImageProposalCreatedEvent.Execute,
	)
}

