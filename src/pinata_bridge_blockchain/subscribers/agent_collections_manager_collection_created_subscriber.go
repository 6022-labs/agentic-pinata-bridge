package subscribers

import (
	"context"
	"fmt"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/zap"
)

type AgentCollectionsManagerCollectionCreatedSubscriber struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
	managers         *settings.AgentCollectionsManagersSettings
}

func NewAgentCollectionsManagerCollectionCreatedSubscriber(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
	managers *settings.AgentCollectionsManagersSettings,
) *AgentCollectionsManagerCollectionCreatedSubscriber {
	return &AgentCollectionsManagerCollectionCreatedSubscriber{
		logger:           logger,
		ethClientFactory: ethClientFactory,
		managers:         managers,
	}
}

func (s *AgentCollectionsManagerCollectionCreatedSubscriber) SubscribeCollectionCreated(ctx context.Context, chainId uint64, logs chan<- *abi.AgentCollectionsManagerCollectionCreated) (ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, err
	}

	managerAddress, ok := s.managers.Get(chainId)
	if !ok {
		return nil, fmt.Errorf("no agent collections manager configured for chain %d", chainId)
	}

	agentCollectionsManager, err := abi.NewAgentCollectionsManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	sub, err := agentCollectionsManager.WatchCollectionCreated(nil, logs, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollectionsManager.CollectionCreated events",
		zap.Uint64("chain_id", chainId),
		zap.String("address", managerAddress.Hex()),
	)

	return sub, nil
}
