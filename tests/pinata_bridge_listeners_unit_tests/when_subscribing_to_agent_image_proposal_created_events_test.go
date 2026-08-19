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
	metrics_mocks_pin "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	interfaces_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

const agentImageProposalCreatedEventName = "AgentCollection.AgentImageProposalCreated"

type WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite struct {
	sut *pinata_bridge_listeners.AgentCollectionAgentImageProposalCreatedListener

	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	subscriptionProvider             *interfaces_mocks.MockAgentImageProposalCreatedSubscriptionProviderInterface
	chainEventMetrics                *metrics_mocks.MockChainEventMetricsInterface
}

func WhenSubscribingToAgentImageProposalCreatedEventsBeforeEach(
	t *testing.T,
) *WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite {
	mockController := gomock.NewController(t)

	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
		mockController,
	)
	subscriptionProvider := interfaces_mocks.NewMockAgentImageProposalCreatedSubscriptionProviderInterface(
		mockController,
	)
	chainEventMetrics := metrics_mocks.NewMockChainEventMetricsInterface(mockController)

	// No event reaches the handler in these subscription tests; it only has to be wired.
	handleEvent := use_cases.NewHandleAgentImageProposalCreatedEvent(use_cases.NewPushImageOfAgentImageProposal(
		zap.NewNop(),
		interfaces_mocks.NewMockCidPinnerInterface(mockController),
		interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController),
		metrics_mocks_pin.NewMockPinMetricsInterface(mockController),
	))

	sut := pinata_bridge_listeners.NewAgentCollectionAgentImageProposalCreatedListener(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		use_cases.NewListCollectionAddresses(agentCollectionsManagerRequester),
		chainEventMetrics,
		subscriptionProvider,
		handleEvent,
	)

	return &WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite{
		sut: sut,

		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		subscriptionProvider:             subscriptionProvider,
		chainEventMetrics:                chainEventMetrics,
	}
}

func TestWhenSubscribingToAgentImageProposalCreatedEvents(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	t.Run("Given the chain has one collection", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{collectionAddress}, nil)
			suite.subscriptionProvider.EXPECT().
				StartAgentImageProposalCreatedSubscription(gomock.Any(), testChainId, []common.Address{collectionAddress}).
				Return(make(chan *abi.AgentCollectionV1AgentImageProposalCreated), newStubSubscription(), nil)
		}

		t.Run("Should subscribe to it", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageProposalCreatedEventsBeforeEach(t)

			initSuite(suite)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), agentImageProposalCreatedEventName, testChainId)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), agentImageProposalCreatedEventName, testChainId).AnyTimes()

			err := suite.sut.SubscribeAll(context.Background())

			assert.NoError(t, err)
		})
	})

	t.Run("Given the collection cannot be listed", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return(nil, assert.AnError)
		}

		t.Run("Should return the error without subscribing", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageProposalCreatedEventsBeforeEach(t)

			initSuite(suite)

			err := suite.sut.SubscribeAll(context.Background())

			assert.Equal(t, assert.AnError, err)
		})
	})

	t.Run("Given a subscription that unsubscribes cleanly", func(t *testing.T) {
		t.Parallel()

		subscription := newStubSubscription()
		stoppedTracking := make(chan struct{})

		initSuite := func(suite *WhenSubscribingToAgentImageProposalCreatedEventsTestingSuite) {
			suite.subscriptionProvider.EXPECT().
				StartAgentImageProposalCreatedSubscription(gomock.Any(), testChainId, []common.Address{collectionAddress}).
				Return(make(chan *abi.AgentCollectionV1AgentImageProposalCreated), subscription, nil)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), agentImageProposalCreatedEventName, testChainId)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), agentImageProposalCreatedEventName, testChainId).
				Do(func(context.Context, string, uint64) { close(stoppedTracking) })
		}

		t.Run("Should record the subscription as closed once its watcher stops", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToAgentImageProposalCreatedEventsBeforeEach(t)
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
