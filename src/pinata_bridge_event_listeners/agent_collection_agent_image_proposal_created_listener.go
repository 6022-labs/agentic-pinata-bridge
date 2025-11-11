package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/subscribers"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageProposalCreatedListener struct {
	logger                                             *zap.Logger
	agentImageProposalCreatedEventHandler              event_handlers.AgentImageProposalCreatedEventHandlerInterface
	agentCollectionAgentImageProposalCreatedSubscriber subscribers.AgentCollectionAgentImageProposalCreatedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan *abi.AgentCollectionV1AgentImageProposalCreated
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionAgentImageProposalCreatedListener(
	logger *zap.Logger,
	agentImageProposalCreatedEventHandler event_handlers.AgentImageProposalCreatedEventHandlerInterface,
	agentCollectionAgentImageProposalCreatedSubscriber subscribers.AgentCollectionAgentImageProposalCreatedSubscriberInterface,
) *AgentCollectionAgentImageProposalCreatedListener {
	return &AgentCollectionAgentImageProposalCreatedListener{
		logger:                                logger,
		agentImageProposalCreatedEventHandler: agentImageProposalCreatedEventHandler,
		agentCollectionAgentImageProposalCreatedSubscriber: agentCollectionAgentImageProposalCreatedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan *abi.AgentCollectionV1AgentImageProposalCreated),
	}
}

func (listener *AgentCollectionAgentImageProposalCreatedListener) Listen() error {
	for {
		select {
		case event := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.AgentImageProposalCreated event", zap.Any("event", event))

			err := listener.agentImageProposalCreatedEventHandler.Handle(event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.AgentImageProposalCreated event", zap.Any("event", event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionAgentImageProposalCreatedListener) Subscribe(collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.AgentImageProposalCreated events", zap.String("collectionAddress", collectionAddress.Hex()))

	subscription, err := listener.agentCollectionAgentImageProposalCreatedSubscriber.SubscribeAgentImageProposalCreated(ctx, collectionAddress, listener.eventChannel)
	if err != nil {
		return err
	}

	go func() {
		if err := <-subscription.Err(); err != nil {
			listener.errorChannel <- err
		}
	}()

	listener.logger.Info("Listening for AgentCollection.AgentImageProposalCreated events", zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
