package pinata_bridge_listeners_unit_tests

import (
	"context"
	"testing"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_listeners_mocks/metrics_mocks/interfaces_mocks"
	interfaces_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

const agentImageUpdatedEventName = "AgentCollection.AgentImageUpdated"

type WhenSubscribingToAgentImageUpdatedEventsTestingSuite struct {
	sut *pinata_bridge_listeners.AgentCollectionAgentImageUpdatedListener

	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	subscriptionProvider             *interfaces_mocks.MockAgentImageUpdatedSubscriptionProviderInterface
	chainEventMetrics                *metrics_mocks.MockChainEventMetricsInterface
}

func WhenSubscribingToAgentImageUpdatedEventsBeforeEach(
	t *testing.T,
) *WhenSubscribingToAgentImageUpdatedEventsTestingSuite {
	mockController := gomock.NewController(t)

	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
		mockController,
	)
	subscriptionProvider := interfaces_mocks.NewMockAgentImageUpdatedSubscriptionProviderInterface(
		mockController,
	)
	chainEventMetrics := metrics_mocks.NewMockChainEventMetricsInterface(mockController)

	// No event reaches the handler in these subscription tests; it only has to be wired.
	handleEvent := interfaces_mocks.NewMockAgentImageUpdatedEventHandlerInterface(mockController)

	sut := pinata_bridge_listeners.NewAgentCollectionAgentImageUpdatedListener(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		use_cases.NewListCollectionAddresses(agentCollectionsManagerRequester),
		chainEventMetrics,
		subscriptionProvider,
		handleEvent,
	)

	return &WhenSubscribingToAgentImageUpdatedEventsTestingSuite{
		sut: sut,

		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		subscriptionProvider:             subscriptionProvider,
		chainEventMetrics:                chainEventMetrics,
	}
}

func TestWhenSubscribingToAgentImageUpdatedEvents(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	t.Run("Given the chain has one collection", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToAgentImageUpdatedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{collectionAddress}, nil)
			suite.subscriptionProvider.EXPECT().
				StartAgentImageUpdatedSubscription(gomock.Any(), testChainId, []common.Address{collectionAddress}).
				Return(make(chan *abi.AgentCollectionV1AgentImageUpdated), newStubSubscription(), nil)
		}

		t.Run("Should subscribe to it", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageUpdatedEventsBeforeEach(t)

			initSuite(suite)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), agentImageUpdatedEventName, testChainId)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), agentImageUpdatedEventName, testChainId).AnyTimes()

			err := suite.sut.SubscribeAll(context.Background())

			assert.NoError(t, err)
		})
	})

	t.Run("Given the collection cannot be listed", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToAgentImageUpdatedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return(nil, assert.AnError)
		}

		t.Run("Should return the error without subscribing", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageUpdatedEventsBeforeEach(t)

			initSuite(suite)

			err := suite.sut.SubscribeAll(context.Background())

			assert.Equal(t, assert.AnError, err)
		})
	})

	t.Run("Given a subscription that unsubscribes cleanly", func(t *testing.T) {
		t.Parallel()

		subscription := newStubSubscription()
		stoppedTracking := make(chan struct{})

		initSuite := func(suite *WhenSubscribingToAgentImageUpdatedEventsTestingSuite) {
			suite.subscriptionProvider.EXPECT().
				StartAgentImageUpdatedSubscription(gomock.Any(), testChainId, []common.Address{collectionAddress}).
				Return(make(chan *abi.AgentCollectionV1AgentImageUpdated), subscription, nil)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), agentImageUpdatedEventName, testChainId)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), agentImageUpdatedEventName, testChainId).
				Do(func(context.Context, string, uint64) { close(stoppedTracking) })
		}

		t.Run("Should record the subscription as closed once its watcher stops", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageUpdatedEventsBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Subscribe(context.Background(), testChainId, collectionAddress)
			assert.NoError(t, err)

			// A nil error is a clean unsubscribe; the watcher must still report the subscription closed.
			subscription.errors <- nil

			select {
			case <-stoppedTracking:
			case <-time.After(time.Second):
				assert.Fail(t, "the watcher stopped without recording the subscription as closed")
			}
		})
	})
}
