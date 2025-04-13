package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/event_subscribers"
	"go.uber.org/zap"
)

type AgenticAIAgentCollectionMintedListener struct {
	logger                                   *zap.Logger
	mintedEventHandler                       event_handlers.MintedEventHandlerInterface
	agenticAIAgentCollectionMintedSubscriber event_subscribers.AgenticAIAgentCollectionMintedSubscriberInterface
}

func NewAgenticAIAgentCollectionMintedListener(
	logger *zap.Logger,
	mintedEventHandler event_handlers.MintedEventHandlerInterface,
	agenticAIAgentCollectionMintedSubscriber event_subscribers.AgenticAIAgentCollectionMintedSubscriberInterface,
) *AgenticAIAgentCollectionMintedListener {
	return &AgenticAIAgentCollectionMintedListener{
		logger:                                   logger,
		mintedEventHandler:                       mintedEventHandler,
		agenticAIAgentCollectionMintedSubscriber: agenticAIAgentCollectionMintedSubscriber,
	}
}

func (listener *AgenticAIAgentCollectionMintedListener) Listen() error {
	ctx := context.Background()

	listener.logger.Info("Starting AgenticAIAgentCollectionMintedListener")

	events := make(chan *abi.AgenticAIAgentCollectionMinted)

	subscription, err := listener.agenticAIAgentCollectionMintedSubscriber.SubscribeMinted(ctx, events)
	if err != nil {
		return err
	}

	listener.logger.Debug("Subscribed to AgenticAIAgentCollectionMinted")

	for {
		select {
		case event := <-events:
			listener.logger.Info("Received AgenticAIAgentCollectionMinted event", zap.Any("event", event))

			err = listener.mintedEventHandler.Handle(event)

			if err != nil {
				listener.logger.Error("failed to finish agent creation", zap.Any("event", event), zap.Error(err))
			}
		case err := <-subscription.Err():
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}
