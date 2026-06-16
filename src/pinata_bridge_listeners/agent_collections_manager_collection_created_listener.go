package pinata_bridge_listeners

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionsManagerCollectionCreatedListener struct {
	logger                     *zap.Logger
	chainsSettings             *settings.ChainsSettings
	collectionEventSubscribers []CollectionEventSubscriberInterface
	agentCollectionsManagerCollectionCreatedEventSubscriptionProvider interfaces.AgentCollectionsManagerCollectionCreatedEventSubscriptionProviderInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]
	subscriptions []ethereum.Subscription
}

type newAgentCollectionsManagerCollectionCreatedListenerParams struct {
	dig.In

	Logger                     *zap.Logger
	ChainsSettings             *settings.ChainsSettings
	CollectionEventSubscribers []CollectionEventSubscriberInterface `group:"collection_event_subscribers"`
	AgentCollectionsManagerCollectionCreatedEventSubscriptionProvider interfaces.AgentCollectionsManagerCollectionCreatedEventSubscriptionProviderInterface
}

func NewAgentCollectionsManagerCollectionCreatedListener(
	params newAgentCollectionsManagerCollectionCreatedListenerParams,
) *AgentCollectionsManagerCollectionCreatedListener {
	return &AgentCollectionsManagerCollectionCreatedListener{
		logger:                     params.Logger,
		chainsSettings:             params.ChainsSettings,
		collectionEventSubscribers: params.CollectionEventSubscribers,
		agentCollectionsManagerCollectionCreatedEventSubscriptionProvider: params.AgentCollectionsManagerCollectionCreatedEventSubscriptionProvider,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]),
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) SubscribeAll() error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		if err := listener.subscribe(chainId); err != nil {
			return err
		}
	}

	return nil
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) Listen() error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollectionsManager.CollectionCreated event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			for _, collectionEventSubscriber := range listener.collectionEventSubscribers {
				if err := collectionEventSubscriber.Subscribe(received.chainId, received.event.CollectionAddress); err != nil {
					listener.logger.Error("Failed to notify collection event subscriber about AgentCollectionsManager.CollectionCreated event", zap.Any("event", received.event), zap.Error(err))
				}
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) subscribe(chainId uint64) error {
	listener.logger.Debug("Subscribing to AgentCollectionsManager.CollectionCreated events", zap.Uint64("chainId", chainId))

	rawEvents, subscription, err := listener.agentCollectionsManagerCollectionCreatedEventSubscriptionProvider.StartCollectionCreatedSubscription(chainId)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionsManagerCollectionCreated]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- err
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollectionsManager.CollectionCreated events", zap.Uint64("chainId", chainId))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
