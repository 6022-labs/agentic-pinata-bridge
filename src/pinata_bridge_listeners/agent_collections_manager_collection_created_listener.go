package pinata_bridge_listeners

import (
	"context"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

const collectionCreatedEventName = "AgentCollectionsManager.CollectionCreated"

// AgentCollectionsManagerCollectionCreatedListener watches the manager contract and tells every
// per-collection listener to start watching a collection as soon as it is created.
type AgentCollectionsManagerCollectionCreatedListener struct {
	*AbstractUpdatableSubscriptionListener

	logger                                *zap.Logger
	chainsSettings                        *settings.ChainsSettings
	collectionEventSubscribers            []CollectionEventSubscriberInterface
	chainEventMetrics                     metrics_interfaces.ChainEventMetricsInterface
	collectionCreatedSubscriptionProvider interfaces.CollectionCreatedSubscriptionProviderInterface

	subscribeContext context.Context
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
	listener := &AgentCollectionsManagerCollectionCreatedListener{
		logger:                                params.Logger,
		chainsSettings:                        params.ChainsSettings,
		collectionEventSubscribers:            params.CollectionEventSubscribers,
		chainEventMetrics:                     params.ChainEventMetrics,
		collectionCreatedSubscriptionProvider: params.CollectionCreatedSubscriptionProvider,
		subscribeContext:                      context.Background(),
	}
	listener.AbstractUpdatableSubscriptionListener = NewAbstractUpdatableSubscriptionListener(
		listener.rebuildSubscription,
	)

	return listener
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) SubscribeAll(ctx context.Context) error {
	listener.mutex.Lock()
	listener.subscribeContext = ctx
	listener.mutex.Unlock()

	for _, chainId := range listener.chainsSettings.ChainIds() {
		if err := listener.rebuildSubscription(chainId); err != nil {
			return err
		}
	}

	return nil
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) Listen(ctx context.Context) error {
	var err error

	select {
	case received := <-listener.errorChannel:
		err = received.err
		listener.chainEventMetrics.RecordSubscriptionError(ctx, collectionCreatedEventName, received.chainId)
		listener.logger.Error("Subscription error",
			zap.Uint64("chainId", received.chainId),
			zap.Error(err),
		)
	case <-ctx.Done():
	}

	listener.stop()

	return err
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) rebuildSubscription(chainId uint64) error {
	listener.mutex.Lock()
	ctx := listener.subscribeContext
	listener.mutex.Unlock()

	rawEvents, subscription, err := listener.collectionCreatedSubscriptionProvider.StartCollectionCreatedSubscription(
		ctx,
		chainId,
	)
	if err != nil {
		return err
	}

	listener.replaceSubscription(chainId, subscription)
	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, collectionCreatedEventName, chainId)
	listener.startWatcher(ctx, chainId, rawEvents, subscription)

	listener.logger.Info("Watching the collections manager", zap.Uint64("chainId", chainId))

	return nil
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) startWatcher(
	ctx context.Context,
	chainId uint64,
	rawEvents <-chan *abi.AgentCollectionsManagerCollectionCreated,
	subscription ethereum.Subscription,
) {
	listener.waitGroup.Add(1)

	go func() {
		defer listener.waitGroup.Done()
		defer listener.chainEventMetrics.RecordSubscriptionClosed(ctx, collectionCreatedEventName, chainId)

		for {
			select {
			case event, open := <-rawEvents:
				if !open {
					return
				}
				if event == nil {
					continue
				}

				listener.dispatch(ctx, chainId, event)
			case err := <-subscription.Err():
				if err == nil {
					return
				}
				listener.reportError(chainId, err)

				return
			}
		}
	}()
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) dispatch(
	ctx context.Context,
	chainId uint64,
	event *abi.AgentCollectionsManagerCollectionCreated,
) {
	listener.logger.Info("Received "+collectionCreatedEventName+" event",
		zap.Uint64("chainId", chainId),
		zap.Any("event", event),
	)

	start := time.Now()
	outcome := metrics_interfaces.ChainEventOutcomeHandled
	for _, collectionEventSubscriber := range listener.collectionEventSubscribers {
		if err := collectionEventSubscriber.Subscribe(ctx, chainId, event.CollectionAddress); err != nil {
			listener.logger.Error("Failed to watch the created collection",
				zap.Any("event", event),
				zap.Error(err),
			)
			outcome = metrics_interfaces.ChainEventOutcomeFailed
		}
	}
	listener.chainEventMetrics.RecordEvent(ctx, collectionCreatedEventName, chainId, outcome, time.Since(start))

	listener.onBlockSeen(chainId, event.RawLog().BlockNumber)
}
