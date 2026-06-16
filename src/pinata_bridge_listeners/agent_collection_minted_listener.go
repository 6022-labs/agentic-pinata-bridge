package pinata_bridge_listeners

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintedListener struct {
	logger                                         *zap.Logger
	chainsSettings                                 *settings.ChainsSettings
	mintedEventHandler                             interfaces.MintedEventHandlerInterface
	agentCollectionsManagerRequester               interfaces.AgentCollectionsManagerRequesterInterface
	agentCollectionMintedEventSubscriptionProvider interfaces.AgentCollectionMintedEventSubscriptionProviderInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1Minted]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	mintedEventHandler interfaces.MintedEventHandlerInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	agentCollectionMintedEventSubscriptionProvider interfaces.AgentCollectionMintedEventSubscriptionProviderInterface,
) *AgentCollectionMintedListener {
	return &AgentCollectionMintedListener{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		mintedEventHandler:               mintedEventHandler,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentCollectionMintedEventSubscriptionProvider: agentCollectionMintedEventSubscriptionProvider,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1Minted]),
	}
}

func (listener *AgentCollectionMintedListener) SubscribeAll() error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		collections, err := listener.agentCollectionsManagerRequester.GetAllCollectionAddresses(chainId)
		if err != nil {
			return err
		}

		for _, collection := range collections {
			if err := listener.Subscribe(chainId, collection); err != nil {
				return err
			}
		}
	}

	return nil
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
	listener.logger.Debug("Subscribing to AgentCollection.Minted events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents, subscription, err := listener.agentCollectionMintedEventSubscriptionProvider.StartMintedSubscription(chainId, collectionAddress)
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
