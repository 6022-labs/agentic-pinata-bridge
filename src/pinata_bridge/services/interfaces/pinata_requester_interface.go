package interfaces

import "context"

type PinataRequesterInterface interface {
	PinCid(ctx context.Context, cid string, hostNodes []string) error
	IsCidUploaded(ctx context.Context, cid string) (*bool, error)
}
