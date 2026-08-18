package pinata_bridge_listeners

import (
	"context"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionsManagerCollectionCreatedListener struct {
	logger                                *zap.Logger
	chainsSettings                        *settings.ChainsSettings
	collectionEventSubscribers            []CollectionEventSubscriberInterface
	chainEventMetrics                     metrics_interfaces.ChainEventMetricsInterface
	collectionCreatedSubscriptionProvider interfaces.CollectionCreatedSubscriptionProviderInterface

	errorChannel  chan ChainSubscriptionError
	eventChannel  chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]
	subscriptions []ethereum.Subscription
}

type newAgentCollectionsManagerCollectionCreatedListenerParams struct {
	dig.In

	Logger                                *zap.Logger
	ChainsSettings                        *settings.ChainsSettings
	CollectionEventSubscribers            []CollectionEventSubscriberInterface `group:"collection_event_subscribers"`
	ChainEventMetrics                     metrics_interfaces.ChainEventMetricsInterface
	CollectionCreatedSubscriptionProvider interfaces.CollectionCreatedSubscriptionProviderInterface
}

func NewAgentCollectionsManagerCollectionCreatedListener(
	params newAgentCollectionsManagerCollectionCreatedListenerParams,
) *AgentCollectionsManagerCollectionCreatedListener {
	return &AgentCollectionsManagerCollectionCreatedListener{
		logger:                                params.Logger,
		chainsSettings:                        params.ChainsSettings,
		collectionEventSubscribers:            params.CollectionEventSubscribers,
		chainEventMetrics:                     params.ChainEventMetrics,
		collectionCreatedSubscriptionProvider: params.CollectionCreatedSubscriptionProvider,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan ChainSubscriptionError),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]),
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) SubscribeAll(ctx context.Context) error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		if err := listener.subscribe(ctx, chainId); err != nil {
			return err
		}
	}

	return nil
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) Listen(ctx context.Context) error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info(
				"Received AgentCollectionsManager.CollectionCreated event",
				zap.Uint64("chainId", received.chainId),
				zap.Any("event", received.event),
			)

			start := time.Now()
			outcome := metrics_interfaces.ChainEventOutcomeHandled
			for _, collectionEventSubscriber := range listener.collectionEventSubscribers {
				if err := collectionEventSubscriber.Subscribe(
					ctx,
					received.chainId,
					received.event.CollectionAddress,
				); err != nil {
					listener.logger.Error(
						"Failed to notify collection event subscriber about AgentCollectionsManager.CollectionCreated event",
						zap.Any("event", received.event),
						zap.Error(err),
					)
					outcome = metrics_interfaces.ChainEventOutcomeFailed
				}
			}
			listener.chainEventMetrics.RecordEvent(
				ctx,
				"AgentCollectionsManager.CollectionCreated",
				received.chainId,
				outcome,
				time.Since(start),
			)
		case received := <-listener.errorChannel:
			listener.chainEventMetrics.RecordSubscriptionError(
				ctx,
				"AgentCollectionsManager.CollectionCreated",
				received.chainId,
			)
			listener.logger.Error("Subscription error",
				zap.Uint64("chainId", received.chainId),
				zap.Error(received.err),
			)

			return received.err
		}
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) subscribe(ctx context.Context, chainId uint64) error {
	listener.logger.Debug(
		"Subscribing to AgentCollectionsManager.CollectionCreated events",
		zap.Uint64("chainId", chainId),
	)

	rawEvents, subscription, err := listener.collectionCreatedSubscriptionProvider.StartCollectionCreatedSubscription(
		ctx,
		chainId,
	)
	if err != nil {
		return err
	}

	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, "AgentCollectionsManager.CollectionCreated", chainId)

	go func() {
		defer listener.chainEventMetrics.RecordSubscriptionClosed(
			ctx,
			"AgentCollectionsManager.CollectionCreated",
			chainId,
		)

		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionsManagerCollectionCreated]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- ChainSubscriptionError{chainId: chainId, err: err}
				}
				return
			}
		}
	}()

	listener.logger.Info(
		"Listening for AgentCollectionsManager.CollectionCreated events",
		zap.Uint64("chainId", chainId),
	)
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
