package pinata_bridge_listeners

import (
	"context"
	"sync"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners/metrics/interfaces"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// StartCollectionSubscription opens one chain+collection subscription in the blockchain adapter.
type StartCollectionSubscription[T any] func(
	ctx context.Context,
	chainId uint64,
	collectionAddress common.Address,
) (<-chan *T, ethereum.Subscription, error)

// HandleChainEvent hands one decoded event to the use case that owns it.
type HandleChainEvent[T any] func(ctx context.Context, chainId uint64, event *T) error

// ChainEventListener is the transport shared by every per-collection event listener: it opens
// subscriptions, pumps raw events onto one channel and hands each to a single use case.
type ChainEventListener[T any] struct {
	logger                  *zap.Logger
	eventName               string
	chainsSettings          *settings.ChainsSettings
	listCollectionAddresses *use_cases.ListCollectionAddresses
	chainEventMetrics       metrics_interfaces.ChainEventMetricsInterface
	startSubscription       StartCollectionSubscription[T]
	handleChainEvent        HandleChainEvent[T]

	errorChannel chan ChainSubscriptionError
	eventChannel chan ChainEvent[T]

	// Subscribe is called from the collection-created listener's goroutine as well as SubscribeAll.
	subscriptionsMutex sync.Mutex
	subscriptions      []ethereum.Subscription
}

func NewChainEventListener[T any](
	logger *zap.Logger,
	eventName string,
	chainsSettings *settings.ChainsSettings,
	listCollectionAddresses *use_cases.ListCollectionAddresses,
	chainEventMetrics metrics_interfaces.ChainEventMetricsInterface,
	startSubscription StartCollectionSubscription[T],
	handleChainEvent HandleChainEvent[T],
) *ChainEventListener[T] {
	return &ChainEventListener[T]{
		logger:                  logger,
		eventName:               eventName,
		chainsSettings:          chainsSettings,
		listCollectionAddresses: listCollectionAddresses,
		chainEventMetrics:       chainEventMetrics,
		startSubscription:       startSubscription,
		handleChainEvent:        handleChainEvent,

		subscriptions: []ethereum.Subscription{},
		errorChannel:  make(chan ChainSubscriptionError),
		eventChannel:  make(chan ChainEvent[T]),
	}
}

func (listener *ChainEventListener[T]) SubscribeAll(ctx context.Context) error {
	for _, chainId := range listener.chainsSettings.ChainIds() {
		collections, err := listener.listCollectionAddresses.Execute(ctx, chainId)
		if err != nil {
			return err
		}

		for _, collection := range collections {
			if err := listener.Subscribe(ctx, chainId, collection); err != nil {
				return err
			}
		}
	}

	return nil
}

func (listener *ChainEventListener[T]) Listen(ctx context.Context) error {
	for {
		select {
		case received := <-listener.eventChannel:
			listener.logger.Info("Received "+listener.eventName+" event",
				zap.Uint64("chainId", received.chainId),
				zap.Any("event", received.event),
			)

			start := time.Now()
			outcome := metrics_interfaces.ChainEventOutcomeHandled
			if err := listener.handleChainEvent(ctx, received.chainId, received.event); err != nil {
				listener.logger.Error("Failed to handle "+listener.eventName+" event",
					zap.Any("event", received.event),
					zap.Error(err),
				)
				outcome = metrics_interfaces.ChainEventOutcomeFailed
			}
			listener.chainEventMetrics.RecordEvent(
				ctx,
				listener.eventName,
				received.chainId,
				outcome,
				time.Since(start),
			)
		case received := <-listener.errorChannel:
			listener.chainEventMetrics.RecordSubscriptionError(ctx, listener.eventName, received.chainId)
			listener.logger.Error("Subscription error",
				zap.Uint64("chainId", received.chainId),
				zap.Error(received.err),
			)

			return received.err
		}
	}
}

func (listener *ChainEventListener[T]) Subscribe(
	ctx context.Context,
	chainId uint64,
	collectionAddress common.Address,
) error {
	listener.logger.Debug("Subscribing to "+listener.eventName+" events",
		zap.Uint64("chainId", chainId),
		zap.String("collectionAddress", collectionAddress.Hex()),
	)

	rawEvents, subscription, err := listener.startSubscription(ctx, chainId, collectionAddress)
	if err != nil {
		return err
	}

	listener.chainEventMetrics.RecordSubscriptionOpened(ctx, listener.eventName, chainId)

	go func() {
		defer listener.chainEventMetrics.RecordSubscriptionClosed(ctx, listener.eventName, chainId)

		for {
			select {
			case event := <-rawEvents:
				listener.eventChannel <- ChainEvent[T]{chainId: chainId, event: event}
			case err := <-subscription.Err():
				if err != nil {
					listener.errorChannel <- ChainSubscriptionError{chainId: chainId, err: err}
				}
				return
			}
		}
	}()

	listener.logger.Info("Listening for "+listener.eventName+" events",
		zap.Uint64("chainId", chainId),
		zap.String("collectionAddress", collectionAddress.Hex()),
	)

	listener.subscriptionsMutex.Lock()
	listener.subscriptions = append(listener.subscriptions, subscription)
	listener.subscriptionsMutex.Unlock()

	return nil
}
