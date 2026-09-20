package pinata_bridge_listeners

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"go.uber.org/zap"
)

type (
	AgentCollectionMintedListener              = ChainEventListener[abi.AgentCollectionV1Minted]
	AgentCollectionMintProposalCreatedListener = ChainEventListener[abi.AgentCollectionV1MintProposalCreated]
	// AgentCollectionAgentImageProposalCreatedListener watches proposals that replace an agent's image.
	AgentCollectionAgentImageProposalCreatedListener = ChainEventListener[abi.AgentCollectionV1AgentImageProposalCreated]
	// AgentCollectionAgentImageUpdatedListener catches image writes that skip the proposal flow.
	AgentCollectionAgentImageUpdatedListener = ChainEventListener[abi.AgentCollectionV1AgentImageUpdated]
)

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	mintedSubscriptionProvider interfaces.MintedSubscriptionProviderInterface,
	mintedEventHandler interfaces.MintedEventHandlerInterface,
) *AgentCollectionMintedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.Minted",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		mintedSubscriptionProvider.StartMintedSubscription,
		mintedEventHandler.Handle,
	)
}

func NewAgentCollectionMintProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	mintProposalCreatedSubscriptionProvider interfaces.MintProposalCreatedSubscriptionProviderInterface,
	mintProposalCreatedEventHandler interfaces.MintProposalCreatedEventHandlerInterface,
) *AgentCollectionMintProposalCreatedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.MintProposalCreated",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		mintProposalCreatedSubscriptionProvider.StartMintProposalCreatedSubscription,
		mintProposalCreatedEventHandler.Handle,
	)
}

func NewAgentCollectionAgentImageProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	agentImageProposalCreatedSubscriptionProvider interfaces.AgentImageProposalCreatedSubscriptionProviderInterface,
	agentImageProposalCreatedEventHandler interfaces.AgentImageProposalCreatedEventHandlerInterface,
) *AgentCollectionAgentImageProposalCreatedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.AgentImageProposalCreated",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		agentImageProposalCreatedSubscriptionProvider.StartAgentImageProposalCreatedSubscription,
		agentImageProposalCreatedEventHandler.Handle,
	)
}

func NewAgentCollectionAgentImageUpdatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	agentImageUpdatedSubscriptionProvider interfaces.AgentImageUpdatedSubscriptionProviderInterface,
	agentImageUpdatedEventHandler interfaces.AgentImageUpdatedEventHandlerInterface,
) *AgentCollectionAgentImageUpdatedListener {
	return NewChainEventListener(
		logger,
		"AgentCollection.AgentImageUpdated",
		chainsSettings,
		listCollectionAddresses,
		chainEventMetrics,
		agentImageUpdatedSubscriptionProvider.StartAgentImageUpdatedSubscription,
		agentImageUpdatedEventHandler.Handle,
	)
}
