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
	ipfsCheckRequester                services.IpfsCheckRequesterInterface
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface
}

func NewPushAgentImageCidToPinata(
	logger *zap.Logger,
	pinataRequester services.PinataRequesterInterface,
	ipfsCheckRequester services.IpfsCheckRequesterInterface,
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface,
) *PushAgentImageCidToPinata {
	return &PushAgentImageCidToPinata{
		logger:                            logger,
		pinataRequester:                   pinataRequester,
		ipfsCheckRequester:                ipfsCheckRequester,
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
	addresses, err := p.getCidHostAddresses(cid)
	if err != nil {
		p.logger.Warn("Failed to get host addresses for cid", zap.String("cid", cid), zap.Error(err))
	}
	if len(addresses) == 0 {
		p.logger.Warn("No host addresses found for cid", zap.String("cid", cid))
	}

	err = p.pinataRequester.PinCidToPinata(cid, addresses)
	if err != nil {
		return err
	}
	return nil
}

func (p *PushAgentImageCidToPinata) getCidHostAddresses(cid string) ([]string, error) {
	var addresses []string
	var err error

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		addresses, err = p.ipfsCheckRequester.GetMultiAddresses(cid)
		if err == nil && len(addresses) > 0 {
			return addresses, nil
		}

		p.logger.Warn("Failed to get host addresses for cid, retrying...", zap.String("cid", cid), zap.Int("attempt", attempt), zap.Error(err))
	}

	return nil, err
}
