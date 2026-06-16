package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionsManagerCollectionCreatedListener struct {
	logger                                             *zap.Logger
	chainsSettings                                     *settings.ChainsSettings
	collectionListeners                                []CollectionListenerInterface
	agentCollectionsManagerCollectionCreatedSubscriber interfaces.AgentCollectionsManagerCollectionCreatedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]
	subscriptions []ethereum.Subscription
}

type newAgentCollectionsManagerCollectionCreatedListenerParams struct {
	dig.In

	Logger                                             *zap.Logger
	ChainsSettings                                     *settings.ChainsSettings
	CollectionListeners                                []CollectionListenerInterface `group:"collection_listeners"`
	AgentCollectionsManagerCollectionCreatedSubscriber interfaces.AgentCollectionsManagerCollectionCreatedSubscriberInterface
}

func NewAgentCollectionsManagerCollectionCreatedListener(
	params newAgentCollectionsManagerCollectionCreatedListenerParams,
) *AgentCollectionsManagerCollectionCreatedListener {
	return &AgentCollectionsManagerCollectionCreatedListener{
		logger:              params.Logger,
		chainsSettings:      params.ChainsSettings,
		collectionListeners: params.CollectionListeners,
		agentCollectionsManagerCollectionCreatedSubscriber: params.AgentCollectionsManagerCollectionCreatedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionsManagerCollectionCreated]),
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) Listen() error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		if err := listener.subscribe(chainId); err != nil {
			return err
		}
	}

	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollectionsManager.CollectionCreated event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			for _, collectionListener := range listener.collectionListeners {
				if err := collectionListener.Subscribe(received.chainId, received.event.CollectionAddress); err != nil {
					listener.logger.Error("Failed to notify collection listener about AgentCollectionsManager.CollectionCreated event", zap.Any("event", received.event), zap.Error(err))
				}
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionsManagerCollectionCreatedListener) subscribe(chainId uint64) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollectionsManager.CollectionCreated events", zap.Uint64("chainId", chainId))

	rawEvents := make(chan *abi.AgentCollectionsManagerCollectionCreated)
	subscription, err := listener.agentCollectionsManagerCollectionCreatedSubscriber.SubscribeCollectionCreated(ctx, chainId, rawEvents)
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
