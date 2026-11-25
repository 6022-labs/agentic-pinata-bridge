package services

import (
	"context"
	"errors"
	"io"
	"net/url"
	"time"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/models"
	"github.com/avast/retry-go/v5"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	checkRetryAttempts = 3
	checkRetryDelay    = time.Second
)

type IpfsCheckRequester struct {
	logger            *zap.Logger
	ipfsCheckClient   clients.IpfsCheckClientInterface
	hostLookupMetrics metrics_interfaces.HostLookupMetricsInterface
}

func NewIpfsCheckRequester(
	logger *zap.Logger,
	ipfsCheckClient clients.IpfsCheckClientInterface,
	hostLookupMetrics metrics_interfaces.HostLookupMetricsInterface,
) *IpfsCheckRequester {
	return &IpfsCheckRequester{
		logger:            logger,
		ipfsCheckClient:   ipfsCheckClient,
		hostLookupMetrics: hostLookupMetrics,
	}
}

func (r *IpfsCheckRequester) GetHostNodeIds(ctx context.Context, cid string) ([]string, error) {
	checkResponses, attempts, err := r.checkWithRetry(ctx, cid)
	if err != nil {
		r.hostLookupMetrics.RecordHostLookup(ctx, metrics_interfaces.HostLookupOutcomeFailed, attempts)
		return nil, err
	}

	hostNodeIds := make([]string, 0)
	for _, provider := range checkResponses {
		// An empty ConnectionMaddrs means ipfs-check never reached the provider; Pinata gains nothing from it.
		if provider.ID == "" || len(provider.ConnectionMaddrs) == 0 {
			continue
		}

		hostNodeIds = append(hostNodeIds, provider.ID)
	}

	outcome := metrics_interfaces.HostLookupOutcomeFound
	if len(hostNodeIds) == 0 {
		outcome = metrics_interfaces.HostLookupOutcomeEmpty
	}
	r.hostLookupMetrics.RecordHostLookup(ctx, outcome, attempts)

	// ipfs-check reports a provider once per source that advertised it.
	return lo.Uniq(hostNodeIds), nil
}

func (r *IpfsCheckRequester) checkWithRetry(
	ctx context.Context,
	cid string,
) ([]models.ExternalCheckResponse, int64, error) {
	attempts := int64(0)
	checkResponses, err := retry.NewWithData[[]models.ExternalCheckResponse](
		retry.Context(ctx),
		retry.Attempts(checkRetryAttempts),
		retry.Delay(checkRetryDelay),
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)),
		retry.LastErrorOnly(true),
		retry.RetryIf(r.isRetryable),
		retry.OnRetry(func(_ uint, err error) {
			r.logger.Warn("Failed to get host node ids for cid, retrying...", zap.String("cid", cid), zap.Error(err))
		}),
	).Do(func() ([]models.ExternalCheckResponse, error) {
		attempts++
		return r.ipfsCheckClient.Check(ctx, cid)
	})

	return checkResponses, attempts, err
}

// isRetryable leaves timeouts out, as the pin goes ahead without host nodes anyway.
func (r *IpfsCheckRequester) isRetryable(err error) bool {
	var statusError *apperrors.UpstreamStatusError
	if errors.As(err, &statusError) {
		return statusError.IsRetryable()
	}

	var transportError *url.Error
	if errors.As(err, &transportError) {
		return !transportError.Timeout() && !errors.Is(err, context.Canceled)
	}

	return errors.Is(err, io.ErrUnexpectedEOF)
}
