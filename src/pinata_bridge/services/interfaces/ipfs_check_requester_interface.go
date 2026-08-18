package interfaces

import "context"

type IpfsCheckRequesterInterface interface {
	GetMultiAddresses(ctx context.Context, cid string) ([]string, error)
}
