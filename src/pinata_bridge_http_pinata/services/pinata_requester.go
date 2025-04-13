package services

import (
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/models"
)

type PinataRequester struct {
	client clients.PinataClientInterface
}

func NewPinataRequester(client clients.PinataClientInterface) *PinataRequester {
	return &PinataRequester{
		client: client,
	}
}

func (pr *PinataRequester) PinCidToPinata(cid string, hostNodes []string) error {
	request := &models.ExternalPinByHashRequest{
		HashToPin: cid,
	}

	if len(hostNodes) > 0 {
		request.PinataOptions = &models.ExternalPinataOptions{
			HostNodes: hostNodes,
		}
	}

	_, err := pr.client.PinByHash(request)
	if err != nil {
		return err
	}

	return nil
}
