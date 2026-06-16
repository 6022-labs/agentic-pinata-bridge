package pinata_bridge_event_listeners

// ChainEvent pairs a decoded event with the chain it was observed on.
type ChainEvent[T any] struct {
	chainId uint64
	event   *T
}
