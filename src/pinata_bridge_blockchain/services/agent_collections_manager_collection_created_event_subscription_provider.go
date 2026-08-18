package services

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"fmt"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum"
	"go.uber.org/zap"
)

type CollectionCreatedSubscriptionProvider struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
	managers         *settings.AgentCollectionsManagersSettings
}

func NewCollectionCreatedSubscriptionProvider(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
	managers *settings.AgentCollectionsManagersSettings,
) *CollectionCreatedSubscriptionProvider {
	return &CollectionCreatedSubscriptionProvider{
		logger:           logger,
		ethClientFactory: ethClientFactory,
		managers:         managers,
	}
}

func (s *CollectionCreatedSubscriptionProvider) StartCollectionCreatedSubscription(ctx context.Context, chainId uint64) (<-chan *abi.AgentCollectionsManagerCollectionCreated, ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, nil, err
	}

	managerAddress, ok := s.managers.Get(chainId)
	if !ok {
		return nil, nil, fmt.Errorf("no agent collections manager configured for chain %d", chainId)
	}

	agentCollectionsManager, err := abi.NewAgentCollectionsManager(managerAddress, client)
	if err != nil {
		return nil, nil, err
	}

	logs := make(chan *abi.AgentCollectionsManagerCollectionCreated, 64)
	sub, err := agentCollectionsManager.WatchCollectionCreated(&bind.WatchOpts{Context: ctx}, logs, nil)
	if err != nil {
		return nil, nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollectionsManager.CollectionCreated events",
		zap.Uint64("chain_id", chainId),
		zap.String("address", managerAddress.Hex()),
	)

	return logs, sub, nil
}
