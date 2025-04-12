package use_cases

import (
	"math/big"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/services"
	"go.uber.org/zap"
)

type PushAgentImageCidToPinataInterface interface {
	PushAllAgentImageCids() error
	PushFromAgentTokenId(tokenId big.Int) error
	PushFromAgentImageCid(cid string) error
}

type PushAgentImageCidToPinata struct {
	logger                            *zap.Logger
	pinataRequester                   services.PinataRequesterInterface
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface
}

func NewPushAgentImageCidToPinata(
	logger *zap.Logger,
	pinataRequester services.PinataRequesterInterface,
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface,
) *PushAgentImageCidToPinata {
	return &PushAgentImageCidToPinata{
		logger:                            logger,
		pinataRequester:                   pinataRequester,
		agenticAIAgentCollectionRequester: agenticAIAgentCollectionRequester,
	}
}

func (p *PushAgentImageCidToPinata) PushAllAgentImageCids() error {
	tokenIds, err := p.agenticAIAgentCollectionRequester.GetAllTokenIds()
	if err != nil {
		return err
	}

	for _, tokenId := range tokenIds {
		err = p.PushFromAgentTokenId(tokenId)
		if err != nil {
			p.logger.Error("Failed to push agent image cid to pinata", zap.Error(err))
			continue
		}

		p.logger.Info("Successfully pushed agent image cid to pinata", zap.String("tokenId", tokenId.String()))
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushFromAgentTokenId(tokenId big.Int) error {
	cid, err := p.agenticAIAgentCollectionRequester.GetAgentImage(tokenId)
	if err != nil {
		return err
	}

	err = p.PushFromAgentImageCid(*cid)
	if err != nil {
		return err
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushFromAgentImageCid(cid string) error {
	err := p.pinataRequester.PinCidToPinata(cid)
	if err != nil {
		return err
	}
	return nil
}
