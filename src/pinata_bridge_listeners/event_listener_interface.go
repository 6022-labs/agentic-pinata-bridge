package pinata_bridge_listeners

import "github.com/ethereum/go-ethereum/common"

type EventListenerInterface interface {
	SubscribeAll() error
	Listen() error
}

type CollectionEventSubscriberInterface interface {
	Subscribe(chainId uint64, collectionAddress common.Address) error
}
