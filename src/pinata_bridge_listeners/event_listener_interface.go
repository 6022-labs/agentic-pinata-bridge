package pinata_bridge_listeners

import "context"

// EventListenerInterface is one chain-event listener the host starts and runs.
type EventListenerInterface interface {
	// SubscribeAll opens a subscription per configured chain before Listen is called.
	SubscribeAll(ctx context.Context) error
	Listen(ctx context.Context) error
}
