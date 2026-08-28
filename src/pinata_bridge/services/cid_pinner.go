package services

import (
	"context"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	traces_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces/interfaces"
	"go.uber.org/zap"
)

const hostNodeIdsMaxRetries = 3

// CidPinner pins a single CID, looking its host addresses up first so Pinata can fetch it faster.
type CidPinner struct {
	logger             *zap.Logger
	pinataRequester    interfaces.PinataRequesterInterface
	ipfsCheckRequester interfaces.IpfsCheckRequesterInterface
	pinMetrics         metrics_interfaces.PinMetricsInterface
	pinTracer          traces_interfaces.PinTracerInterface
}

func NewCidPinner(
	logger *zap.Logger,
	pinataRequester interfaces.PinataRequesterInterface,
	ipfsCheckRequester interfaces.IpfsCheckRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
	pinTracer traces_interfaces.PinTracerInterface,
) *CidPinner {
	return &CidPinner{
		logger:             logger,
		pinataRequester:    pinataRequester,
		ipfsCheckRequester: ipfsCheckRequester,
		pinMetrics:         pinMetrics,
		pinTracer:          pinTracer,
	}
}

func (s *CidPinner) Pin(ctx context.Context, cid string) (err error) {
	ctx, span := s.pinTracer.StartPin(ctx, cid)
	defer span.End()
	defer func() {
		if err != nil {
			span.Fail(err)
		}
	}()

	hostNodeIds, err := s.getCidHostNodeIds(ctx, cid)
	if err != nil {
		s.logger.Warn("Failed to get host node ids for cid", zap.String("cid", cid), zap.Error(err))
	}
	if len(hostNodeIds) == 0 {
		s.logger.Warn("No host node ids found for cid", zap.String("cid", cid))
	}

	withHostAddresses := len(hostNodeIds) > 0

	start := time.Now()
	err = s.pinataRequester.PinCid(ctx, cid, hostNodeIds)
	if err == nil {
		s.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomePinned, withHostAddresses, time.Since(start))
		return nil
	}
	s.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomeFailed, withHostAddresses, time.Since(start))

	// Retrying is only worth it when the failed attempt actually carried host addresses to drop.
	if !withHostAddresses {
		return err
	}

	s.logger.Warn("Failed to pin cid to pinata with host node ids, retrying without",
		zap.String("cid", cid),
		zap.Error(err),
	)

	start = time.Now()
	err = s.pinataRequester.PinCid(ctx, cid, nil)
	if err != nil {
		s.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomeFailed, false, time.Since(start))
		return err
	}

	s.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomePinned, false, time.Since(start))
	return nil
}

func (s *CidPinner) getCidHostNodeIds(ctx context.Context, cid string) ([]string, error) {
	var hostNodeIds []string
	var err error

	for attempt := 1; attempt <= hostNodeIdsMaxRetries; attempt++ {
		hostNodeIds, err = s.ipfsCheckRequester.GetHostNodeIds(ctx, cid)
		if err == nil && len(hostNodeIds) > 0 {
			s.pinMetrics.RecordHostLookup(ctx, metrics_interfaces.HostLookupOutcomeFound, int64(attempt))
			return hostNodeIds, nil
		}

		s.logger.Warn("Failed to get host node ids for cid, retrying...",
			zap.String("cid", cid),
			zap.Int("attempt", attempt),
			zap.Error(err),
		)
	}

	outcome := metrics_interfaces.HostLookupOutcomeEmpty
	if err != nil {
		outcome = metrics_interfaces.HostLookupOutcomeFailed
	}
	s.pinMetrics.RecordHostLookup(ctx, outcome, hostNodeIdsMaxRetries)

	return nil, err
}
