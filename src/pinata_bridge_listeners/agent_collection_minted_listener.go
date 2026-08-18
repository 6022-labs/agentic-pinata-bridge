package pinata_bridge_listeners

import (
	"context"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintedListener struct {
	logger                                         *zap.Logger
	chainsSettings                                 *settings.ChainsSettings
	mintedEventHandler                             interfaces.MintedEventHandlerInterface
	agentCollectionsManagerRequester               interfaces.AgentCollectionsManagerRequesterInterface
	chainEventMetrics                              metrics_interfaces.ChainEventMetricsInterface
	agentCollectionMintedEventSubscriptionProvider interfaces.AgentCollectionMintedEventSubscriptionProviderInterface

	errorChannel  chan ChainSubscriptionError
	eventChannel  chan ChainEvent[abi.AgentCollectionV1Minted]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	mintedEventHandler interfaces.MintedEventHandlerInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	agentCollectionMintedEventSubscriptionProvider interfaces.AgentCollectionMintedEventSubscriptionProviderInterface,
) *AgentCollectionMintedListener {
	return &AgentCollectionMintedListener{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		mintedEventHandler:               mintedEventHandler,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentCollectionMintedEventSubscriptionProvider: agentCollectionMintedEventSubscriptionProvider,
		chainEventMetrics: chainEventMetrics,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan ChainSubscriptionError),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1Minted]),
	}
}

func (listener *AgentCollectionMintedListener) SubscribeAll(ctx context.Context) error {
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

func (listener *AgentCollectionMintedListener) Listen(ctx context.Context) error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.Minted event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			start := time.Now()
			err := listener.mintedEventHandler.Handle(ctx, received.chainId, received.event)
			outcome := metrics_interfaces.ChainEventOutcomeHandled
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.Minted event", zap.Any("event", received.event), zap.Error(err))
				outcome = metrics_interfaces.ChainEventOutcomeFailed
			}
			listener.chainEventMetrics.RecordEvent(ctx, "AgentCollection.Minted", received.chainId, outcome, time.Since(start))
		case received := <-listener.errorChannel:
			listener.chainEventMetrics.RecordSubscriptionError(ctx, "AgentCollection.Minted", received.chainId)
			listener.logger.Error("Subscription error",
				zap.Uint64("chainId", received.chainId),
				zap.Error(received.err),
			)

			return received.err
		}
	}
}

func (listener *AgentCollectionMintedListener) Subscribe(ctx context.Context, chainId uint64, collectionAddress common.Address) error {
	listener.logger.Debug("Subscribing to AgentCollection.Minted events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents, subscription, err := listener.agentCollectionMintedEventSubscriptionProvider.StartMintedSubscription(ctx, chainId, collectionAddress)
	if err != nil {
		return err
	}

	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, "AgentCollection.Minted", chainId)

	go func() {
		defer listener.chainEventMetrics.RecordSubscriptionClosed(ctx, "AgentCollection.Minted", chainId)

		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionV1Minted]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- ChainSubscriptionError{chainId: chainId, err: err}
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollection.Minted events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
