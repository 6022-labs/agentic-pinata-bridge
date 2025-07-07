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

type AgentCollectionMintedSubscriber struct {
	logger *zap.Logger
	client *ethclient.Client
}

type newAgentCollectionMintedSubscriberParams struct {
	dig.In

	Logger *zap.Logger
	Client *ethclient.Client `name:"ws"`
}

func NewAgentCollectionMintedSubscriber(
	params newAgentCollectionMintedSubscriberParams,
) *AgentCollectionMintedSubscriber {
	return &AgentCollectionMintedSubscriber{
		logger: params.Logger,
		client: params.Client,
	}
}

func (s *AgentCollectionMintedSubscriber) SubscribeMinted(ctx context.Context, agentCollectionAddress common.Address, logs chan<- *abi.AgentCollectionV1Minted) (ethereum.Subscription, error) {
	agentCollection, err := abi.NewAgentCollectionV1(agentCollectionAddress, s.client)
	if err != nil {
		return nil, err
	}

	sub, err := agentCollection.WatchMinted(nil, logs, nil, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgentCollection.Minted events",
		zap.String("address", agentCollectionAddress.Hex()),
	)

	return sub, nil
}
