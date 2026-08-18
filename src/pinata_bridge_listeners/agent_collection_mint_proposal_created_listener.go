package pinata_bridge_listeners

import (
	"context"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedListener struct {
	logger                                                      *zap.Logger
	chainsSettings                                              *settings.ChainsSettings
	mintProposalCreatedEventHandler                             interfaces.MintProposalCreatedEventHandlerInterface
	agentCollectionsManagerRequester                            interfaces.AgentCollectionsManagerRequesterInterface
	chainEventMetrics                                           metrics_interfaces.ChainEventMetricsInterface
	agentCollectionMintProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface

	errorChannel  chan ChainSubscriptionError
	eventChannel  chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	mintProposalCreatedEventHandler interfaces.MintProposalCreatedEventHandlerInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	agentCollectionMintProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface,
) *AgentCollectionMintProposalCreatedListener {
	return &AgentCollectionMintProposalCreatedListener{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		mintProposalCreatedEventHandler:  mintProposalCreatedEventHandler,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentCollectionMintProposalCreatedEventSubscriptionProvider: agentCollectionMintProposalCreatedEventSubscriptionProvider,
		chainEventMetrics: chainEventMetrics,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan ChainSubscriptionError),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]),
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) SubscribeAll(ctx context.Context) error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		collections, err := listener.agentCollectionsManagerRequester.GetAllCollectionAddresses(ctx, chainId)
		if err != nil {
			return err
		}

		for _, collection := range collections {
			if err := listener.Subscribe(ctx, chainId, collection); err != nil {
				return err
			}
		}
	}

	return nil
}

func (listener *AgentCollectionMintProposalCreatedListener) Listen(ctx context.Context) error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.MintProposalCreated event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			start := time.Now()
			err := listener.mintProposalCreatedEventHandler.Handle(ctx, received.chainId, received.event)
			outcome := metrics_interfaces.ChainEventOutcomeHandled
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.MintProposalCreated event", zap.Any("event", received.event), zap.Error(err))
				outcome = metrics_interfaces.ChainEventOutcomeFailed
			}
			listener.chainEventMetrics.RecordEvent(ctx, "AgentCollection.MintProposalCreated", received.chainId, outcome, time.Since(start))
		case received := <-listener.errorChannel:
			listener.chainEventMetrics.RecordSubscriptionError(ctx, "AgentCollection.MintProposalCreated", received.chainId)
			listener.logger.Error("Subscription error",
				zap.Uint64("chainId", received.chainId),
				zap.Error(received.err),
			)

			return received.err
		}
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) Subscribe(ctx context.Context, chainId uint64, collectionAddress common.Address) error {
	listener.logger.Debug("Subscribing to AgentCollection.MintProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents, subscription, err := listener.agentCollectionMintProposalCreatedEventSubscriptionProvider.StartMintProposalCreatedSubscription(ctx, chainId, collectionAddress)
	if err != nil {
		return err
	}

	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, "AgentCollection.MintProposalCreated", chainId)

	go func() {
		defer listener.chainEventMetrics.RecordSubscriptionClosed(ctx, "AgentCollection.MintProposalCreated", chainId)

		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionV1MintProposalCreated]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- ChainSubscriptionError{chainId: chainId, err: err}
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollection.MintProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
