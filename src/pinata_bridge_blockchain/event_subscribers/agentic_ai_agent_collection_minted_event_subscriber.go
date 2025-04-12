package event_subscribers

import (
	"context"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_blockchain/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AgenticAIAgentCollectionMintedSubscriber struct {
	logger                           *zap.Logger
	client                           *ethclient.Client
	agenticAIAgentCollectionSettings *settings.AgenticAIAgentCollectionSettings
}

type newAgenticAIAgentCollectionMintedSubscriberParams struct {
	dig.In

	Logger                           *zap.Logger
	Client                           *ethclient.Client `name:"ws"`
	AgenticAIAgentCollectionSettings *settings.AgenticAIAgentCollectionSettings
}

func NewAgenticAIAgentCollectionMintedSubscriber(
	params newAgenticAIAgentCollectionMintedSubscriberParams,
) *AgenticAIAgentCollectionMintedSubscriber {
	return &AgenticAIAgentCollectionMintedSubscriber{
		logger:                           params.Logger,
		client:                           params.Client,
		agenticAIAgentCollectionSettings: params.AgenticAIAgentCollectionSettings,
	}
}

func (s *AgenticAIAgentCollectionMintedSubscriber) SubscribeMinted(ctx context.Context, logs chan<- *abi.AgenticAIAgentCollectionMinted) (ethereum.Subscription, error) {
	agenticAIAgentCollection, err := abi.NewAgenticAIAgentCollection(s.agenticAIAgentCollectionSettings.SmartContractAddress, s.client)
	if err != nil {
		return nil, err
	}

	sub, err := agenticAIAgentCollection.WatchMinted(nil, logs, nil, nil)
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"Subscribed to AgenticAIAgentCollection Minted events",
		zap.String("address", s.agenticAIAgentCollectionSettings.SmartContractAddress.Hex()),
	)

	return sub, nil
}
