package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageUpdatedEventSubscriptionProvider struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
}

func NewAgentCollectionAgentImageUpdatedEventSubscriptionProvider(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
) *AgentCollectionAgentImageUpdatedEventSubscriptionProvider {
	return &AgentCollectionAgentImageUpdatedEventSubscriptionProvider{
		logger:           logger,
		ethClientFactory: ethClientFactory,
	}
}

func (s *AgentCollectionAgentImageUpdatedEventSubscriptionProvider) StartAgentImageUpdatedSubscription(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddresses []common.Address,
) (<-chan *abi.AgentCollectionV1AgentImageUpdated, ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, nil, err
	}

	contractAbi, err := abi.AgentCollectionV1MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	filterer, err := abi.NewAgentCollectionV1Filterer(common.Address{}, client)
	if err != nil {
		return nil, nil, err
	}

	logs, subscription, err := subscribeToCollectionEvent(
		ctx,
		client,
		contractAbi,
		"AgentImageUpdated",
		agentCollectionAddresses,
		func(rawLog types.Log) (*abi.AgentCollectionV1AgentImageUpdated, error) {
			return filterer.ParseAgentImageUpdated(rawLog)
		},
	)
	if err != nil {
		return nil, nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.AgentImageUpdated events",
		zap.Uint64("chainId", chainId),
		zap.Int("collectionCount", len(agentCollectionAddresses)),
	)

	return logs, subscription, nil
}
