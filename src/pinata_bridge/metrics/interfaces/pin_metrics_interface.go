package interfaces

import (
	"context"
	"time"
)

const (
	PinOutcomePinned        = "pinned"
	PinOutcomeAlreadyPinned = "already_pinned"
	PinOutcomeFailed        = "failed"
	PinOutcomeInvalidCid    = "invalid_cid"
)

const (
	SweepKindAll           = "all"
	SweepKindAgent         = "agent"
	SweepKindMintProposal  = "mint_proposal"
	SweepKindImageProposal = "image_proposal"
)

// PinMetricsInterface records the pinata.* metrics of the CID pinning pipeline.
type PinMetricsInterface interface {
	// RecordPin records one pin attempt; withHostAddresses is false for the retry that drops the host list.
	RecordPin(ctx context.Context, outcome string, withHostAddresses bool, duration time.Duration)
	RecordSweep(ctx context.Context, kind string, duration time.Duration, failed bool)
	// RecordSweepImage counts one image (CID) a sweep visited; the outer all-sweep defers to the per-agent sweeps.
	RecordSweepImage(ctx context.Context, kind, outcome string)
}
