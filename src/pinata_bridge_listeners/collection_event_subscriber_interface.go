package pinata_bridge_listeners

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// CollectionEventSubscriberInterface is a listener that can start watching a collection discovered at runtime.
type CollectionEventSubscriberInterface interface {
	Subscribe(ctx context.Context, chainId uint64, collectionAddress common.Address) error
}
