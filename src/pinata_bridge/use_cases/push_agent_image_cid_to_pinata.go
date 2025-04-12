package use_cases

import (
	"math/big"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/services"
)

type PushAgentImageCidToPinataInterface interface {
	PushFromAgentTokenId(tokenId big.Int) error
	PushFromAgentImageCid(cid string) error
}

type PushAgentImageCidToPinata struct {
	pinataRequester                   services.PinataRequesterInterface
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface
}

func NewPushAgentImageCidToPinata(
	pinataRequester services.PinataRequesterInterface,
	agenticAIAgentCollectionRequester services.AgenticAIAgentCollectionRequesterInterface,
) *PushAgentImageCidToPinata {
	return &PushAgentImageCidToPinata{
		pinataRequester:                   pinataRequester,
		agenticAIAgentCollectionRequester: agenticAIAgentCollectionRequester,
	}
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
