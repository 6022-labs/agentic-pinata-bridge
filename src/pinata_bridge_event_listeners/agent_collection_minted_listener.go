package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/subscribers"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintedListener struct {
	logger                          *zap.Logger
	mintedEventHandler              event_handlers.MintedEventHandlerInterface
	agentCollectionMintedSubscriber subscribers.AgentCollectionMintedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan *abi.AgentCollectionV1Minted
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	mintedEventHandler event_handlers.MintedEventHandlerInterface,
	agentCollectionMintedSubscriber subscribers.AgentCollectionMintedSubscriberInterface,
) *AgentCollectionMintedListener {
	return &AgentCollectionMintedListener{
		logger:                          logger,
		mintedEventHandler:              mintedEventHandler,
		agentCollectionMintedSubscriber: agentCollectionMintedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan *abi.AgentCollectionV1Minted),
	}
}

func (listener *AgentCollectionMintedListener) Listen() error {
	for {
		select {
		case event := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.Minted event", zap.Any("event", event))

			err := listener.mintedEventHandler.Handle(event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.Minted event", zap.Any("event", event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionMintedListener) Subscribe(collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.Minted events", zap.String("collectionAddress", collectionAddress.Hex()))

	subscription, err := listener.agentCollectionMintedSubscriber.SubscribeMinted(ctx, collectionAddress, listener.eventChannel)
	if err != nil {
		return err
	}

	go func() {
		if err := <-subscription.Err(); err != nil {
			listener.errorChannel <- err
		}
	}()

	listener.logger.Info("Listening for AgentCollection.Minted events", zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
