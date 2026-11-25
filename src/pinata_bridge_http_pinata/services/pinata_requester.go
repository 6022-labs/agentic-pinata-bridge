package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/avast/retry-go/v5"
)

const (
	rateLimitAttempts     = 3
	rateLimitDefaultDelay = time.Second
	// Past this, the next push of missing images retries instead of this one waiting.
	rateLimitMaxDelay = 10 * time.Second
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

	return pr.retryRateLimited(ctx, func(ctx context.Context) error {
		response, err := pr.client.PinByCidWithResponse(ctx, body)
		if err != nil {
			return fmt.Errorf("failed to send request to Pinata: %w", err)
		}

		return pr.ensureSuccess(response.HTTPResponse, response.Body)
	})
}

func (pr *PinataRequester) IsCidUploaded(ctx context.Context, cid string) (*bool, error) {
	var response *clients.ListFilesResponse
	err := pr.retryRateLimited(ctx, func(ctx context.Context) error {
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

// retryRateLimited leaves any other failure alone: Pinata may already have pinned the cid.
func (pr *PinataRequester) retryRateLimited(ctx context.Context, request func(ctx context.Context) error) error {
	return retry.New(
		retry.Context(ctx),
		retry.Attempts(rateLimitAttempts),
		retry.DelayType(func(_ uint, err error, _ retry.DelayContext) time.Duration {
			return pr.rateLimitDelay(err)
		}),
		retry.LastErrorOnly(true),
		retry.RetryIf(func(err error) bool {
			var statusError *apperrors.UpstreamStatusError
			return errors.As(err, &statusError) && statusError.IsRateLimited() &&
				statusError.RetryAfter <= rateLimitMaxDelay
		}),
	).Do(func() error {
		requestCtx, cancel := context.WithTimeout(ctx, pr.pinataSettings.RequestTimeout)
		defer cancel()

		return request(requestCtx)
	})
}

func (pr *PinataRequester) rateLimitDelay(err error) time.Duration {
	var statusError *apperrors.UpstreamStatusError
	if errors.As(err, &statusError) && statusError.RetryAfter > 0 {
		return statusError.RetryAfter
	}
	return rateLimitDefaultDelay
}

// ensureSuccess keeps the body in the error; the calling use case is what logs it.
func (pr *PinataRequester) ensureSuccess(httpResponse *http.Response, body []byte) error {
	if httpResponse == nil {
		return fmt.Errorf("pinata returned no response")
	}

	if httpResponse.StatusCode >= http.StatusOK && httpResponse.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	return apperrors.NewUpstreamStatusError(
		"pinata",
		httpResponse.StatusCode,
		pr.retryAfter(httpResponse.Header),
		string(body),
	)
}

// retryAfter reads both forms, seconds and HTTP date; 0 when absent, invalid or past.
func (pr *PinataRequester) retryAfter(header http.Header) time.Duration {
	value := header.Get("Retry-After")

	// Out of range still parses to the highest value, so a huge delay never passes for an absent one.
	seconds, err := strconv.ParseUint(value, 10, 16)
	if err == nil || errors.Is(err, strconv.ErrRange) {
		return time.Duration(seconds) * time.Second
	}

	date, err := http.ParseTime(value)
	if err != nil {
		return 0
	}
	return max(time.Until(date), 0)
}
