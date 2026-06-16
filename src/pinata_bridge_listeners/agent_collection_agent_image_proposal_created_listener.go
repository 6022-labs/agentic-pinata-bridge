package pinata_bridge_listeners

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageProposalCreatedListener struct {
	logger                                                            *zap.Logger
	chainsSettings                                                    *settings.ChainsSettings
	agentImageProposalCreatedEventHandler                             interfaces.AgentImageProposalCreatedEventHandlerInterface
	agentCollectionsManagerRequester                                  interfaces.AgentCollectionsManagerRequesterInterface
	agentCollectionAgentImageProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionAgentImageProposalCreatedEventSubscriptionProviderInterface

	errorChannel  chan error
	eventChannel  chan ChainEvent[abi.AgentCollectionV1AgentImageProposalCreated]
	subscriptions []ethereum.Subscription
}

func NewAgentCollectionAgentImageProposalCreatedListener(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	agentImageProposalCreatedEventHandler interfaces.AgentImageProposalCreatedEventHandlerInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	agentCollectionAgentImageProposalCreatedEventSubscriptionProvider interfaces.AgentCollectionAgentImageProposalCreatedEventSubscriptionProviderInterface,
) *AgentCollectionAgentImageProposalCreatedListener {
	return &AgentCollectionAgentImageProposalCreatedListener{
		logger:                                logger,
		chainsSettings:                        chainsSettings,
		agentImageProposalCreatedEventHandler: agentImageProposalCreatedEventHandler,
		agentCollectionsManagerRequester:      agentCollectionsManagerRequester,
		agentCollectionAgentImageProposalCreatedEventSubscriptionProvider: agentCollectionAgentImageProposalCreatedEventSubscriptionProvider,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan error),
		eventChannel:  make(chan ChainEvent[abi.AgentCollectionV1AgentImageProposalCreated]),
	}
}

func (listener *AgentCollectionAgentImageProposalCreatedListener) SubscribeAll() error {
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
	listener.logger.Debug("Subscribing to AgentCollection.AgentImageProposalCreated events", zap.Uint64("chainId", chainId), zap.String("collectionAddress", collectionAddress.Hex()))

	rawEvents, subscription, err := listener.agentCollectionAgentImageProposalCreatedEventSubscriptionProvider.StartAgentImageProposalCreatedSubscription(chainId, collectionAddress)
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
