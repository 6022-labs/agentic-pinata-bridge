package subscribers

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionMintProposalCreatedSubscriber struct {
	logger *zap.Logger
	client *ethclient.Client
}

type newAgentCollectionMintProposalCreatedSubscriberParams struct {
	dig.In

	Logger *zap.Logger
	Client *ethclient.Client `name:"ws"`
}

func NewAgentCollectionMintProposalCreatedSubscriber(
	params newAgentCollectionMintProposalCreatedSubscriberParams,
) *AgentCollectionMintProposalCreatedSubscriber {
	return &AgentCollectionMintProposalCreatedSubscriber{
		logger: params.Logger,
		client: params.Client,
	}
}

func (s *AgentCollectionMintProposalCreatedSubscriber) SubscribeMintProposalCreated(ctx context.Context, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1MintProposalCreated) (ethereum.Subscription, error) {
	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, s.client)
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
