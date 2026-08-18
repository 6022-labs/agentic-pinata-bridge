package pinata_bridge_listeners

// ChainEvent pairs a decoded event with the chain it was observed on.
type ChainEvent[T any] struct {
	chainId uint64
	event   *T
}

// ChainSubscriptionError pairs a failed subscription with the chain it was watching.
type ChainSubscriptionError struct {
	chainId uint64
	err     error
}
