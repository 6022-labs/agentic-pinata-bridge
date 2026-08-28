package interfaces

import "context"

type IpfsCheckRequesterInterface interface {
	// GetHostNodeIds returns the peer ids of the providers ipfs-check actually reached for the cid.
	GetHostNodeIds(ctx context.Context, cid string) ([]string, error)
}
