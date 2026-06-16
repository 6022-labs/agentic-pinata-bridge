package pinata_bridge_event_listeners

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintedListener struct {
	logger                          *zap.Logger
	mintedEventHandler              interfaces.MintedEventHandlerInterface
	agentCollectionMintedSubscriber interfaces.AgentCollectionMintedSubscriberInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1Minted]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	mintedEventHandler interfaces.MintedEventHandlerInterface,
	agentCollectionMintedSubscriber interfaces.AgentCollectionMintedSubscriberInterface,
) *AgentCollectionMintedListener {
	return &AgentCollectionMintedListener{
		logger:                          logger,
		mintedEventHandler:              mintedEventHandler,
		agentCollectionMintedSubscriber: agentCollectionMintedSubscriber,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1Minted]),
	}
}

func (listener *AgentCollectionMintedListener) Listen() error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received AgentCollection.Minted event", zap.Uint64("chainId", received.chainId), zap.Any("event", received.event))

			err := listener.mintedEventHandler.Handle(received.chainId, received.event)
			if err != nil {
				listener.logger.Error("Failed to handle AgentCollection.Minted event", zap.Any("event", received.event), zap.Error(err))
			}
		case err := <-listener.errorChannel:
			listener.logger.Error("Subscription error", zap.Error(err))

			return err
		}
	}
}

func (listener *AgentCollectionMintedListener) Subscribe(chainId uint64, collectionAddress common.Address) error {
	ctx := context.Background()

	listener.logger.Debug("Subscribing to AgentCollection.Minted events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents := make(chan *abi.AgentCollectionV1Minted)
	subscription, err := listener.agentCollectionMintedSubscriber.SubscribeMinted(ctx, chainId, collectionAddress, rawEvents)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[abi.AgentCollectionV1Minted]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- err
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for AgentCollection.Minted events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))
	listener.subscriptions = append(listener.subscriptions, subscription)

	return nil
}
