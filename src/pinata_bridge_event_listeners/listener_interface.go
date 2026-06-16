package pinata_bridge_event_listeners

import "github.com/ethereum/go-ethereum/common"

type ListenerInterface interface {
	Listen() error
}

type CollectionListenerInterface interface {
	ListenerInterface
	Subscribe(chainId uint64, collectionAddress common.Address) error
}
