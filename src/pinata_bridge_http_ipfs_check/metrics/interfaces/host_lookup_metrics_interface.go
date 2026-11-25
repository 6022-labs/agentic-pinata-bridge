package interfaces

import "context"

const (
	HostLookupOutcomeFound  = "found"
	HostLookupOutcomeEmpty  = "empty"
	HostLookupOutcomeFailed = "failed"
)

type HostLookupMetricsInterface interface {
	// RecordHostLookup records one lookup, including how many attempts it cost.
	RecordHostLookup(ctx context.Context, outcome string, attempts int64)
}
