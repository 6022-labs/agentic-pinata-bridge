package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/avast/retry-go/v5"
)

const (
	pinataRetryAttempts = 3
	pinataRetryDelay    = time.Second
)

type PinataRequester struct {
	client         clients.ClientWithResponsesInterface
	pinataSettings *settings.PinataSettings
}

func NewPinataRequester(
	client clients.ClientWithResponsesInterface,
	pinataSettings *settings.PinataSettings,
) *PinataRequester {
	return &PinataRequester{
		client:         client,
		pinataSettings: pinataSettings,
	}
}

func (pr *PinataRequester) PinCid(ctx context.Context, cid string, hostNodes []string) error {
	body := clients.PinByCidJSONRequestBody{Cid: cid}
	if len(hostNodes) > 0 {
		body.HostNodes = &hostNodes
	}

	return pr.retry(ctx, func(ctx context.Context) error {
		response, err := pr.client.PinByCidWithResponse(ctx, body)
		if err != nil {
			return fmt.Errorf("failed to send request to Pinata: %w", err)
		}

		return pr.ensureSuccess(response.HTTPResponse, response.Body)
	})
}

func (pr *PinataRequester) IsCidUploaded(ctx context.Context, cid string) (*bool, error) {
	var response *clients.ListFilesResponse
	err := pr.retry(ctx, func(ctx context.Context) error {
		var err error
		response, err = pr.client.ListFilesWithResponse(
			ctx,
			clients.Public,
			&clients.ListFilesParams{Cid: &cid},
		)
		if err != nil {
			return fmt.Errorf("failed to send request to Pinata: %w", err)
		}

		return pr.ensureSuccess(response.HTTPResponse, response.Body)
	})
	if err != nil {
		return nil, err
	}

	exist := response.JSON200 != nil &&
		response.JSON200.Data != nil &&
		response.JSON200.Data.Files != nil &&
		len(*response.JSON200.Data.Files) == 1

	return &exist, nil
}

// retry never repeats a timeout: a hung Pinata would hold a sweep several times longer.
func (pr *PinataRequester) retry(ctx context.Context, call func(ctx context.Context) error) error {
	return retry.New(
		retry.Context(ctx),
		retry.Attempts(pinataRetryAttempts),
		retry.Delay(pinataRetryDelay),
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)),
		retry.LastErrorOnly(true),
		retry.RetryIf(func(err error) bool {
			var statusError *apperrors.UpstreamStatusError
			return errors.As(err, &statusError) && statusError.IsRetryable()
		}),
	).Do(func() error {
		requestCtx, cancel := context.WithTimeout(ctx, pr.pinataSettings.RequestTimeout)
		defer cancel()

		return call(requestCtx)
	})
}

// ensureSuccess keeps the body in the error; the calling use case is what logs it.
func (pr *PinataRequester) ensureSuccess(httpResponse *http.Response, body []byte) error {
	if httpResponse == nil {
		return fmt.Errorf("pinata returned no response")
	}

	if httpResponse.StatusCode >= http.StatusOK && httpResponse.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	return apperrors.NewUpstreamStatusError("pinata", httpResponse.StatusCode, string(body))
}
