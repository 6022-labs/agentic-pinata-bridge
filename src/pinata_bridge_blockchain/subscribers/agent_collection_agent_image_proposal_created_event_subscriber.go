package subscribers

import (
	"context"

	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgentCollectionAgentImageProposalCreatedSubscriber struct {
	logger *zap.Logger
	client *ethclient.Client
}

type newAgentCollectionAgentImageProposalCreatedSubscriberParams struct {
	dig.In

	Logger *zap.Logger
	Client *ethclient.Client `name:"ws"`
}

func NewAgentCollectionAgentImageProposalCreatedSubscriber(
	params newAgentCollectionAgentImageProposalCreatedSubscriberParams,
) *AgentCollectionAgentImageProposalCreatedSubscriber {
	return &AgentCollectionAgentImageProposalCreatedSubscriber{
		logger: params.Logger,
		client: params.Client,
	}
}

func (s *AgentCollectionAgentImageProposalCreatedSubscriber) SubscribeAgentImageProposalCreated(ctx context.Context, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1AgentImageProposalCreated) (ethereum.Subscription, error) {
	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, s.client)
	if err != nil {
		return nil, err
	}

	sub, err := agentCollection.WatchAgentImageProposalCreated(nil, logs, nil, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.AgentImageProposalCreated events",
		zap.String("address", agentCollectionAddress.Hex()),
	)

	return sub, nil
}
