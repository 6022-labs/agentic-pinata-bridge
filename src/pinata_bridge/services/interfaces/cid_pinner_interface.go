package interfaces

import "context"

// CidPinnerInterface pins one CID on Pinata, passing the IPFS host addresses when they can be found.
type CidPinnerInterface interface {
	// Pin retries the lookup, then falls back to pinning without host addresses when the first attempt failed.
	Pin(ctx context.Context, cid string) error
}
