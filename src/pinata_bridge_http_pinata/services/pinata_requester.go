package services

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/models"
)

type PinataRequester struct {
	client clients.PinataClientInterface
}

func NewPinataRequester(client clients.PinataClientInterface) *PinataRequester {
	return &PinataRequester{
		client: client,
	}
}

func (pr *PinataRequester) PinCid(cid string, hostNodes []string) error {
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

func (pr *PinataRequester) IsCidUploaded(cid string) (*bool, error) {
	response, err := pr.client.QueryFileByCid(cid)
	if err != nil {
		return nil, err
	}

	exist := len(response.Data.Files) == 1

	return &exist, nil
}
