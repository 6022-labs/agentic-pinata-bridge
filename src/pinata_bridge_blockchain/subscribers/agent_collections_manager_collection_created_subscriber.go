package subscribers

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionsManagerCollectionCreatedSubscriber struct {
	logger                          *zap.Logger
	client                          *ethclient.Client
	agentCollectionsManagerSettings *settings.AgentCollectionsManagerSettings
}

type newAgentCollectionsManagerCollectionCreatedSubscriberParams struct {
	dig.In

	Logger                          *zap.Logger
	Client                          *ethclient.Client `name:"ws"`
	AgentCollectionsManagerSettings *settings.AgentCollectionsManagerSettings
}

func NewAgentCollectionsManagerCollectionCreatedSubscriber(
	params newAgentCollectionsManagerCollectionCreatedSubscriberParams,
) *AgentCollectionsManagerCollectionCreatedSubscriber {
	return &AgentCollectionsManagerCollectionCreatedSubscriber{
		logger:                          params.Logger,
		client:                          params.Client,
		agentCollectionsManagerSettings: params.AgentCollectionsManagerSettings,
	}
}

func (s *AgentCollectionsManagerCollectionCreatedSubscriber) SubscribeCollectionCreated(ctx context.Context, logs chan<- *abi.AgentCollectionsManagerCollectionCreated) (ethereum.Subscription, error) {
	agentCollectionsManager, err := abi.NewAgentCollectionsManager(s.agentCollectionsManagerSettings.SmartContractAddress, s.client)
	if err != nil {
		return nil, err
	}

	sub, err := agentCollectionsManager.WatchCollectionCreated(nil, logs, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollectionsManager.CollectionCreated events",
		zap.String("address", s.agentCollectionsManagerSettings.SmartContractAddress.Hex()),
	)

	return sub, nil
}
