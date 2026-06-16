package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/subscribers"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedListener struct {
	logger                                       *zap.Logger
	mintProposalCreatedEventHandler              event_handlers.MintProposalCreatedEventHandlerInterface
	agentCollectionMintProposalCreatedSubscriber subscribers.AgentCollectionMintProposalCreatedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]
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
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]),
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) Listen() error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.MintProposalCreated event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			err := listener.mintProposalCreatedEventHandler.Handle(received.chainId, received.event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.MintProposalCreated event", zap.Any("event", received.event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) Subscribe(chainId uint64, collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.MintProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents := make(chan *abi.AgentCollectionV1MintProposalCreated)
	subscription, err := listener.agentCollectionMintProposalCreatedSubscriber.SubscribeMintProposalCreated(ctx, chainId, collectionAddress, rawEvents)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionV1MintProposalCreated]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- err
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollection.MintProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
