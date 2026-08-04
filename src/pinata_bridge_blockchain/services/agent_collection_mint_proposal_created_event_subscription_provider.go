package services

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedEventSubscriptionProvider struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
}

func NewAgentCollectionMintProposalCreatedEventSubscriptionProvider(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
) *AgentCollectionMintProposalCreatedEventSubscriptionProvider {
	return &AgentCollectionMintProposalCreatedEventSubscriptionProvider{
		logger:           logger,
		ethClientFactory: ethClientFactory,
	}
}

func (s *AgentCollectionMintProposalCreatedEventSubscriptionProvider) StartMintProposalCreatedSubscription(chainId uint64, agentCollectionAddress common.Address) (<-chan *abi.AgentCollectionV1MintProposalCreated, ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, nil, err
	}

	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, client)
	if err != nil {
		return nil, nil, err
	}

	logs := make(chan *abi.AgentCollectionV1MintProposalCreated, 64)
	sub, err := agentCollection.WatchMintProposalCreated(nil, logs, nil)
	if err != nil {
		return nil, nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.MintProposalCreated events",
		zap.String("address", agentCollectionAddress.Hex()),
	)

	return logs, sub, nil
}
