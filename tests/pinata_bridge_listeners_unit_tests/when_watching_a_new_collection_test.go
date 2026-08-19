package pinata_bridge_listeners_unit_tests

import (
	"context"
	"testing"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWhenWatchingANewCollection(t *testing.T) {
	t.Parallel()

	firstCollection := common.HexToAddress("0x1111111111111111111111111111111111111111")
	secondCollection := common.HexToAddress("0x2222222222222222222222222222222222222222")

	t.Run("Given a collection is created after start-up", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToMintedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{firstCollection}, nil)

			// One subscription per chain: first covering one collection, then rebuilt covering both.
			suite.subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), testChainId, []common.Address{firstCollection}).
				Return(make(chan *abi.AgentCollectionV1Minted), newStubSubscription(), nil)
			suite.subscriptionProvider.EXPECT().
				StartMintedSubscription(
					gomock.Any(),
					testChainId,
					[]common.Address{firstCollection, secondCollection},
				).
				Return(make(chan *abi.AgentCollectionV1Minted), newStubSubscription(), nil)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), mintedEventName, testChainId).Times(2)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), mintedEventName, testChainId).AnyTimes()
		}

		t.Run("Should rebuild the chain subscription to cover both collections", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintedEventsBeforeEach(t)
			initSuite(suite)

			assert.NoError(t, suite.sut.SubscribeAll(context.Background()))
			assert.NoError(t, suite.sut.Subscribe(context.Background(), testChainId, secondCollection))
		})
	})

	t.Run("Given the same collection is announced twice", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToMintedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{firstCollection}, nil)
			suite.subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), testChainId, []common.Address{firstCollection}).
				Return(make(chan *abi.AgentCollectionV1Minted), newStubSubscription(), nil).
				Times(1)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), mintedEventName, testChainId)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), mintedEventName, testChainId).AnyTimes()
		}

		t.Run("Should not open a second subscription for it", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintedEventsBeforeEach(t)
			initSuite(suite)

			assert.NoError(t, suite.sut.SubscribeAll(context.Background()))

			// Already tracked: the listener must treat this as a no-op rather than rebuilding.
			assert.NoError(t, suite.sut.Subscribe(context.Background(), testChainId, firstCollection))
		})
	})

	t.Run("Given the listening context is cancelled", func(t *testing.T) {
		t.Parallel()

		subscription := newStubSubscription()

		initSuite := func(suite *WhenSubscribingToMintedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{firstCollection}, nil)
			suite.subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), testChainId, []common.Address{firstCollection}).
				Return(make(chan *abi.AgentCollectionV1Minted), subscription, nil)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), mintedEventName, testChainId)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), mintedEventName, testChainId).AnyTimes()
		}

		t.Run("Should unsubscribe every subscription it opened", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintedEventsBeforeEach(t)
			initSuite(suite)

			assert.NoError(t, suite.sut.SubscribeAll(context.Background()))

			ctx, cancel := context.WithCancel(context.Background())
			listenDone := make(chan error, 1)
			go func() { listenDone <- suite.sut.Listen(ctx) }()

			cancel()

			select {
			case err := <-listenDone:
				assert.NoError(t, err)
			case <-time.After(2 * time.Second):
				assert.Fail(t, "Listen did not return after the context was cancelled")
			}

			assert.True(t, subscription.unsubscribed(), "the subscription was never unsubscribed")
		})
	})
}
