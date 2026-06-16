package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageProposalCreatedListener struct {
	logger                                             *zap.Logger
	agentImageProposalCreatedEventHandler              interfaces.AgentImageProposalCreatedEventHandlerInterface
	agentCollectionAgentImageProposalCreatedSubscriber interfaces.AgentCollectionAgentImageProposalCreatedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1AgentImageProposalCreated]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionAgentImageProposalCreatedListener(
	logger *zap.Logger,
	agentImageProposalCreatedEventHandler interfaces.AgentImageProposalCreatedEventHandlerInterface,
	agentCollectionAgentImageProposalCreatedSubscriber interfaces.AgentCollectionAgentImageProposalCreatedSubscriberInterface,
) *AgentCollectionAgentImageProposalCreatedListener {
	return &AgentCollectionAgentImageProposalCreatedListener{
		logger:                                logger,
		agentImageProposalCreatedEventHandler: agentImageProposalCreatedEventHandler,
		agentCollectionAgentImageProposalCreatedSubscriber: agentCollectionAgentImageProposalCreatedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1AgentImageProposalCreated]),
	}
}

func (listener *AgentCollectionAgentImageProposalCreatedListener) Listen() error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.AgentImageProposalCreated event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			err := listener.agentImageProposalCreatedEventHandler.Handle(received.chainId, received.event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.AgentImageProposalCreated event", zap.Any("event", received.event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionAgentImageProposalCreatedListener) Subscribe(chainId uint64, collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.AgentImageProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents := make(chan *abi.AgentCollectionV1AgentImageProposalCreated)
	subscription, err := listener.agentCollectionAgentImageProposalCreatedSubscriber.SubscribeAgentImageProposalCreated(ctx, chainId, collectionAddress, rawEvents)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionV1AgentImageProposalCreated]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- err
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollection.AgentImageProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
