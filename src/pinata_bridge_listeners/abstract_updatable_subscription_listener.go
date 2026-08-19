package pinata_bridge_listeners

import (
	"sync"

	"github.com/ethereum/go-ethereum"
)

// AbstractUpdatableSubscriptionListener owns the subscription lifecycle shared by every listener:
// one live subscription per chain, rebuilt when the watched set changes, retired once the
// replacement has caught up, and unsubscribed on shutdown.
type AbstractUpdatableSubscriptionListener struct {
	rebuildSubscription func(chainId uint64) error

	mutex            sync.Mutex
	subscriptions    map[uint64]ethereum.Subscription
	oldSubscriptions map[uint64][]ethereum.Subscription
	lastBlockSeen    map[uint64]uint64
	needsRebuild     map[uint64]bool

	// Buffered so a failing subscription never blocks its own watcher goroutine.
	errorChannel chan error
	waitGroup    sync.WaitGroup
}

func NewAbstractUpdatableSubscriptionListener(
	rebuildSubscription func(chainId uint64) error,
) *AbstractUpdatableSubscriptionListener {
	return &AbstractUpdatableSubscriptionListener{
		rebuildSubscription: rebuildSubscription,
		subscriptions:       map[uint64]ethereum.Subscription{},
		oldSubscriptions:    map[uint64][]ethereum.Subscription{},
		lastBlockSeen:       map[uint64]uint64{},
		needsRebuild:        map[uint64]bool{},
		errorChannel:        make(chan error, 8),
	}
}

// stop unsubscribes everything this listener opened and waits for the watcher goroutines to finish.
func (l *AbstractUpdatableSubscriptionListener) stop() {
	l.mutex.Lock()
	for chainId, subscription := range l.subscriptions {
		subscription.Unsubscribe()
		delete(l.subscriptions, chainId)
	}
	for chainId, subscriptions := range l.oldSubscriptions {
		for _, subscription := range subscriptions {
			subscription.Unsubscribe()
		}
		delete(l.oldSubscriptions, chainId)
	}
	l.mutex.Unlock()

	l.waitGroup.Wait()
}

func (l *AbstractUpdatableSubscriptionListener) markNeedsRebuild(chainId uint64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.needsRebuild[chainId] = true
}

// onBlockSeen advances the chain's high-water mark and retires subscriptions the
// replacement has already overtaken, so no event is dropped across a rebuild.
func (l *AbstractUpdatableSubscriptionListener) onBlockSeen(chainId, blockNumber uint64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	last := l.lastBlockSeen[chainId]
	if blockNumber > last {
		l.lastBlockSeen[chainId] = blockNumber
	}

	// Retire only this chain's parked subscriptions: chains advance independently, so another
	// chain's block number says nothing about whether this chain's replacement has caught up.
	if len(l.oldSubscriptions[chainId]) > 0 && blockNumber >= last {
		for _, subscription := range l.oldSubscriptions[chainId] {
			subscription.Unsubscribe()
		}
		delete(l.oldSubscriptions, chainId)
	}
}

// replaceSubscription installs the new subscription for a chain and parks the previous one
// for retirement rather than cutting it off immediately.
func (l *AbstractUpdatableSubscriptionListener) replaceSubscription(
	chainId uint64,
	subscription ethereum.Subscription,
) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if previous, ok := l.subscriptions[chainId]; ok && previous != nil {
		l.oldSubscriptions[chainId] = append(l.oldSubscriptions[chainId], previous)
	}
	l.subscriptions[chainId] = subscription
}

func (l *AbstractUpdatableSubscriptionListener) reportError(err error) {
	select {
	case l.errorChannel <- err:
	default:
	}
}
