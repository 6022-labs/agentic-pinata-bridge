package subscribers

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedSubscriber struct {
	logger           *zap.Logger
	ethClientFactory *factory.EthClientFactory
}

func NewAgentCollectionMintProposalCreatedSubscriber(
	logger *zap.Logger,
	ethClientFactory *factory.EthClientFactory,
) *AgentCollectionMintProposalCreatedSubscriber {
	return &AgentCollectionMintProposalCreatedSubscriber{
		logger:           logger,
		ethClientFactory: ethClientFactory,
	}
}

func (s *AgentCollectionMintProposalCreatedSubscriber) SubscribeMintProposalCreated(ctx context.Context, chainId uint64, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1MintProposalCreated) (ethereum.Subscription, error) {
	client, err := s.ethClientFactory.Ws(chainId)
	if err != nil {
		return nil, err
	}

	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, client)
	if err != nil {
		return nil, err
	}

	sub, err := agentCollection.WatchMintProposalCreated(nil, logs, nil, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.MintProposalCreated events",
		zap.String("address", agentCollectionAddress.Hex()),
	)

	return sub, nil
}
