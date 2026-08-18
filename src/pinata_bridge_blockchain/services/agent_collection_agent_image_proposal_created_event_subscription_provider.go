package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageProposalCreatedEventSubscriptionProvider struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
}

func NewAgentCollectionAgentImageProposalCreatedEventSubscriptionProvider(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
) *AgentCollectionAgentImageProposalCreatedEventSubscriptionProvider {
	return &AgentCollectionAgentImageProposalCreatedEventSubscriptionProvider{
		logger:           logger,
		ethClientFactory: ethClientFactory,
	}
}

func (s *AgentCollectionAgentImageProposalCreatedEventSubscriptionProvider) StartAgentImageProposalCreatedSubscription(ctx context.Context, chainId uint64, agentCollectionAddress common.Address) (<-chan *abi.AgentCollectionV1AgentImageProposalCreated, ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, nil, err
	}

	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, client)
	if err != nil {
		return nil, nil, err
	}

	logs := make(chan *abi.AgentCollectionV1AgentImageProposalCreated, 64)
	sub, err := agentCollection.WatchAgentImageProposalCreated(&bind.WatchOpts{Context: ctx}, logs, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.AgentImageProposalCreated events",
		zap.String("address", agentCollectionAddress.Hex()),
	)

	return logs, sub, nil
}
