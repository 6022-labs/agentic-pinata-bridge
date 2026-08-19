package pinata_bridge_listeners

import (
	"context"
	"slices"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// ChainEvent is any decoded contract event; RawLog is generated onto every abi event struct.
type ChainEvent interface {
	RawLog() types.Log
}

// StartCollectionSubscription opens one subscription covering every watched collection on a chain.
type StartCollectionSubscription[T any] func(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddresses []common.Address,
) (<-chan *T, ethereum.Subscription, error)

// HandleChainEvent hands one decoded event to the use case that owns it.
type HandleChainEvent[T any] func(ctx context.Context, chainId uint64, event *T) error

// ChainEventListener is transport only: it keeps one subscription per chain covering every known
// collection, and hands each decoded event to a single use case.
type ChainEventListener[T ChainEvent] struct {
	*AbstractUpdatableSubscriptionListener

	logger                  *zap.Logger
	eventName               string
	chainsSettings          *settings.ChainsSettings
	listCollectionAddresses *use_cases.ListCollectionAddresses
	chainEventMetrics       metrics_interfaces.ChainEventMetricsInterface
	startSubscription       StartCollectionSubscription[T]
	handleChainEvent        HandleChainEvent[T]

	// Guarded by the embedded listener's mutex.
	trackedAddresses map[uint64][]common.Address
	subscribeContext context.Context
}

func NewChainEventListener[T ChainEvent](
	logger *zap.Logger,
	eventName string,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	startSubscription StartCollectionSubscription[T],
	handleChainEvent HandleChainEvent[T],
) *ChainEventListener[T] {
	listener := &ChainEventListener[T]{
		logger:                  logger,
		eventName:               eventName,
		chainsSettings:          chainsSettings,
		listCollectionAddresses: listCollectionAddresses,
		chainEventMetrics:       chainEventMetrics,
		startSubscription:       startSubscription,
		handleChainEvent:        handleChainEvent,
		trackedAddresses:        map[uint64][]common.Address{},
		subscribeContext:        context.Background(),
	}
	listener.AbstractUpdatableSubscriptionListener = NewAbstractUpdatableSubscriptionListener(
		listener.rebuildSubscription,
	)

	return listener
}

func (listener *ChainEventListener[T]) SubscribeAll(ctx context.Context) error {
	listener.mutex.Lock()
	listener.subscribeContext = ctx
	listener.mutex.Unlock()

	for _, chainId := range listener.chainsSettings.ChainIds() {
		collections, err := listener.listCollectionAddresses.Execute(ctx, chainId)
		if err != nil {
			return err
		}

		if len(collections) == 0 {
			listener.logger.Info("No collection to watch yet",
				zap.String("event", listener.eventName),
				zap.Uint64("chainId", chainId),
			)
			continue
		}

		listener.mutex.Lock()
		listener.trackedAddresses[chainId] = slices.Clone(collections)
		listener.mutex.Unlock()

		if err := listener.rebuildSubscription(chainId); err != nil {
			return err
		}
	}

	return nil
}

// Subscribe starts watching a collection discovered after start-up; the chain's single
// subscription is rebuilt to cover it.
func (listener *ChainEventListener[T]) Subscribe(
	ctx context.Context,
	chainId uint64,
	collectionAddress common.Address,
) error {
	listener.mutex.Lock()
	if slices.Contains(listener.trackedAddresses[chainId], collectionAddress) {
		listener.mutex.Unlock()
		return nil
	}
	listener.trackedAddresses[chainId] = append(listener.trackedAddresses[chainId], collectionAddress)
	listener.subscribeContext = ctx
	listener.mutex.Unlock()

	return listener.rebuildSubscription(chainId)
}

func (listener *ChainEventListener[T]) Listen(ctx context.Context) error {
	var err error

	select {
	case received := <-listener.errorChannel:
		err = received.err
		listener.chainEventMetrics.RecordSubscriptionError(ctx, listener.eventName, received.chainId)
		listener.logger.Error("Subscription error",
			zap.String("event", listener.eventName),
			zap.Uint64("chainId", received.chainId),
			zap.Error(err),
		)
	case <-ctx.Done():
	}

	listener.stop()

	return err
}

func (listener *ChainEventListener[T]) rebuildSubscription(chainId uint64) error {
	listener.mutex.Lock()
	addresses := slices.Clone(listener.trackedAddresses[chainId])
	ctx := listener.subscribeContext
	listener.mutex.Unlock()

	if len(addresses) == 0 {
		return nil
	}

	rawEvents, subscription, err := listener.startSubscription(ctx, chainId, addresses)
	if err != nil {
		return err
	}

	listener.replaceSubscription(chainId, subscription)
	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, listener.eventName, chainId)
	listener.startWatcher(ctx, chainId, rawEvents, subscription)

	listener.logger.Info("Watching collections",
		zap.String("event", listener.eventName),
		zap.Uint64("chainId", chainId),
		zap.Int("collectionCount", len(addresses)),
	)

	return nil
}

func (listener *ChainEventListener[T]) startWatcher(
	ctx context.Context,
	chainId uint64,
	rawEvents <-chan *T,
	subscription ethereum.Subscription,
) {
	listener.waitGroup.Add(1)

	go func() {
		defer listener.waitGroup.Done()
		defer listener.chainEventMetrics.RecordSubscriptionClosed(ctx, listener.eventName, chainId)

		for {
			select {
			case event, open := <-rawEvents:
				if !open {
					return
				}
				if event == nil {
					continue
				}

				listener.dispatch(ctx, chainId, event)
			case err := <-subscription.Err():
				if err == nil {
					return
				}
				listener.reportError(chainId, err)

				return
			}
		}
	}()
}

func (listener *ChainEventListener[T]) dispatch(ctx context.Context, chainId uint64, event *T) {
	listener.logger.Info("Received "+listener.eventName+" event",
		zap.Uint64("chainId", chainId),
		zap.Any("event", event),
	)

	start := time.Now()
	outcome := metrics_interfaces.ChainEventOutcomeHandled
	if err := listener.handleChainEvent(ctx, chainId, event); err != nil {
		listener.logger.Error("Failed to handle "+listener.eventName+" event",
			zap.Any("event", event),
			zap.Error(err),
		)
		outcome = metrics_interfaces.ChainEventOutcomeFailed
	}
	listener.chainEventMetrics.RecordEvent(ctx, listener.eventName, chainId, outcome, time.Since(start))

	listener.onBlockSeen(chainId, (*event).RawLog().BlockNumber)
}
