package pinata_bridge_listeners_unit_tests

import "github.com/ethereum/go-ethereum"

// stubSubscription drives a listener's watcher goroutine: push to errors to end it, nil for a clean unsubscribe.
type stubSubscription struct {
	errors chan error
}

func newStubSubscription() *stubSubscription {
	return &stubSubscription{errors: make(chan error, 1)}
}

func (s *stubSubscription) Unsubscribe() {}

func (s *stubSubscription) Err() <-chan error { return s.errors }

var _ ethereum.Subscription = (*stubSubscription)(nil)
