package pinata_bridge_listeners

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedListener struct {
	logger                                                       *zap.Logger
	chainsSettings                                               *settings.ChainsSettings
	mintProposalCreatedEventHandler                              interfaces.MintProposalCreatedEventHandlerInterface
	agentCollectionsManagerRequester                            interfaces.AgentCollectionsManagerRequesterInterface
	agentCollectionMintProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionMintProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	mintProposalCreatedEventHandler interfaces.MintProposalCreatedEventHandlerInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	agentCollectionMintProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionMintProposalCreatedEventSubscriptionProviderInterface,
) *AgentCollectionMintProposalCreatedListener {
	return &AgentCollectionMintProposalCreatedListener{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		mintProposalCreatedEventHandler:  mintProposalCreatedEventHandler,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentCollectionMintProposalCreatedEventSubscriptionProvider: agentCollectionMintProposalCreatedEventSubscriptionProvider,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1MintProposalCreated]),
	}
}

func (listener *AgentCollectionMintProposalCreatedListener) SubscribeAll() error {
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
	listener.logger.Debug("Subscribing to AgentCollection.MintProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents, subscription, err := listener.agentCollectionMintProposalCreatedEventSubscriptionProvider.StartMintProposalCreatedSubscription(chainId, collectionAddress)
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
