package pinata_bridge_listeners_unit_tests

import (
	"sync"

	"github.com/ethereum/go-ethereum"
)

// stubSubscription drives a listener's watcher goroutine: push to errors to end it, nil for a clean unsubscribe.
type stubSubscription struct {
	errors chan error

	mutex           sync.Mutex
	unsubscribeSeen bool
}

func newStubSubscription() *stubSubscription {
	return &stubSubscription{errors: make(chan error, 1)}
}

// Unsubscribe closes the error channel, the way go-ethereum subscriptions signal a clean stop.
func (s *stubSubscription) Unsubscribe() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.unsubscribeSeen {
		return
	}
	s.unsubscribeSeen = true
	close(s.errors)
}

func (s *stubSubscription) Err() <-chan error { return s.errors }

func (s *stubSubscription) unsubscribed() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.unsubscribeSeen
}

var _ ethereum.Subscription = (*stubSubscription)(nil)
