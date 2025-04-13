package services

import "github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/clients"

type PinataRequester struct {
	client clients.PinataClientInterface
}

func NewPinataRequester(client clients.PinataClientInterface) *PinataRequester {
	return &PinataRequester{
		client: client,
	}
}

func (pr *PinataRequester) PinCidToPinata(cid string) error {
	_, err := pr.client.PinByHash(cid)
	if err != nil {
		return err
	}

	return nil
}
