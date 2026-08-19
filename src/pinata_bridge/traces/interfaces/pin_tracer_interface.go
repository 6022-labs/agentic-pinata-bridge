package interfaces

import "context"

// PinTracerInterface opens the spans of one pinning operation.
type PinTracerInterface interface {
	// StartPin opens the span covering a single CID being pinned on Pinata.
	StartPin(ctx context.Context, cid string) (context.Context, Span)
	// StartSweep opens the span covering a sweep; kind is all, agent, mint_proposal or image_proposal.
	StartSweep(ctx context.Context, kind string) (context.Context, Span)
}
