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

type AgentCollectionMintProposalCreatedListener struct {
	logger                                       *zap.Logger
	mintProposalCreatedEventHandler              event_handlers.MintProposalCreatedEventHandlerInterface
	agentCollectionMintProposalCreatedSubscriber subscribers.AgentCollectionMintProposalCreatedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan *abi.AgentCollectionV1MintProposalCreated
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintProposalCreatedListener(
	logger *zap.Logger,
	mintProposalCreatedEventHandler event_handlers.MintProposalCreatedEventHandlerInterface,
	agentCollectionMintProposalCreatedSubscriber subscribers.AgentCollectionMintProposalCreatedSubscriberInterface,
) *AgentCollectionMintProposalCreatedListener {
	return &AgentCollectionMintProposalCreatedListener{
		logger:                          logger,
		mintProposalCreatedEventHandler: mintProposalCreatedEventHandler,
		agentCollectionMintProposalCreatedSubscriber: agentCollectionMintProposalCreatedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan *abi.AgentCollectionV1MintProposalCreated),
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) Listen() error {
	for {
		select {
		case event := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.MintProposalCreated event", zap.Any("event", event))

			err := listener.mintProposalCreatedEventHandler.Handle(event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.MintProposalCreated event", zap.Any("event", event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) Subscribe(collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.MintProposalCreated events", zap.String("collectionAddress", collectionAddress.Hex()))

	subscription, err := listener.agentCollectionMintProposalCreatedSubscriber.SubscribeMintProposalCreated(ctx, collectionAddress, listener.eventChannel)
	if err != nil {
		return err
	}

	go func() {
		if err := <-subscription.Err(); err != nil {
			listener.errorChannel <- err
		}
	}()

	listener.logger.Info("Listening for AgentCollection.MintProposalCreated events", zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
