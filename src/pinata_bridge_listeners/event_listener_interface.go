package pinata_bridge_listeners

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type EventListenerInterface interface {
	SubscribeAll(ctx context.Context) error
	Listen(ctx context.Context) error
}

type CollectionEventSubscriberInterface interface {
	Subscribe(ctx context.Context, chainId uint64, collectionAddress common.Address) error
}
