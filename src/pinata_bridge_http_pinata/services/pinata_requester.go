package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
)

type PinataRequester struct {
	client clients.ClientWithResponsesInterface
}

func NewPinataRequester(client clients.ClientWithResponsesInterface) *PinataRequester {
	return &PinataRequester{
		client: client,
	}
}

func (pr *PinataRequester) PinCid(ctx context.Context, cid string, hostNodes []string) error {
	body := clients.PinByCidJSONRequestBody{Cid: cid}
	if len(hostNodes) > 0 {
		body.HostNodes = &hostNodes
	}

	response, err := pr.client.PinByCidWithResponse(ctx, body)
	if err != nil {
		return fmt.Errorf("failed to send request to Pinata: %w", err)
	}

	return pr.ensureSuccess(response.HTTPResponse, response.Body)
}

func (pr *PinataRequester) IsCidUploaded(ctx context.Context, cid string) (*bool, error) {
	response, err := pr.client.ListFilesWithResponse(
		ctx,
		clients.Public,
		&clients.ListFilesParams{Cid: &cid},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Pinata: %w", err)
	}

	if err := pr.ensureSuccess(response.HTTPResponse, response.Body); err != nil {
		return nil, err
	}

	exist := response.JSON200 != nil &&
		response.JSON200.Data != nil &&
		response.JSON200.Data.Files != nil &&
		len(*response.JSON200.Data.Files) == 1

	return &exist, nil
}

// ensureSuccess keeps the body in the error; the calling use case is what logs it.
func (pr *PinataRequester) ensureSuccess(httpResponse *http.Response, body []byte) error {
	if httpResponse == nil {
		return fmt.Errorf("pinata returned no response")
	}

	if httpResponse.StatusCode >= http.StatusOK && httpResponse.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	return fmt.Errorf("pinata returned %d: %s", httpResponse.StatusCode, string(body))
}
