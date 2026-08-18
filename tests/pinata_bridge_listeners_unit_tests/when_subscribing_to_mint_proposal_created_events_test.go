package pinata_bridge_listeners_unit_tests

import (
	"context"
	"testing"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_blockchain/abi"
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

const mintProposalCreatedEventName = "AgentCollection.MintProposalCreated"

type WhenSubscribingToMintProposalCreatedEventsTestingSuite struct {
	sut *pinata_bridge_listeners.AgentCollectionMintProposalCreatedListener

	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	subscriptionProvider             *interfaces_mocks.MockMintProposalCreatedSubscriptionProviderInterface
	chainEventMetrics                *metrics_mocks.MockChainEventMetricsInterface
}

func WhenSubscribingToMintProposalCreatedEventsBeforeEach(t *testing.T) *WhenSubscribingToMintProposalCreatedEventsTestingSuite {
	mockController := gomock.NewController(t)

	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController)
	subscriptionProvider := interfaces_mocks.NewMockMintProposalCreatedSubscriptionProviderInterface(mockController)
	chainEventMetrics := metrics_mocks.NewMockChainEventMetricsInterface(mockController)

	// No event reaches the handler in these subscription tests; it only has to be wired.
	handleEvent := use_cases.NewHandleMintProposalCreatedEvent(use_cases.NewPushImagesOfMintProposal(
		zap.NewNop(),
		interfaces_mocks.NewMockCidPinnerInterface(mockController),
		interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController),
		metrics_mocks_pin.NewMockPinMetricsInterface(mockController),
	))

	sut := pinata_bridge_listeners.NewAgentCollectionMintProposalCreatedListener(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		use_cases.NewListCollectionAddresses(agentCollectionsManagerRequester),
		chainEventMetrics,
		subscriptionProvider,
		handleEvent,
	)

	return &WhenSubscribingToMintProposalCreatedEventsTestingSuite{
		sut: sut,

		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		subscriptionProvider:             subscriptionProvider,
		chainEventMetrics:                chainEventMetrics,
	}
}

func TestWhenSubscribingToMintProposalCreatedEvents(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	t.Run("Given the chain has one collection", func(t *testing.T) {
		t.Parallel()

		t.Run("Should subscribe to it", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintProposalCreatedEventsBeforeEach(t)

			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{collectionAddress}, nil)
			suite.subscriptionProvider.EXPECT().
				StartMintProposalCreatedSubscription(gomock.Any(), testChainId, collectionAddress).
				Return(make(chan *abi.AgentCollectionV1MintProposalCreated), newStubSubscription(), nil)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), mintProposalCreatedEventName, testChainId)
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), mintProposalCreatedEventName, testChainId).AnyTimes()

			err := suite.sut.SubscribeAll(context.Background())

			assert.NoError(t, err)
		})
	})

	t.Run("Given the collection cannot be listed", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the error without subscribing", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintProposalCreatedEventsBeforeEach(t)

			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return(nil, assert.AnError)

			err := suite.sut.SubscribeAll(context.Background())

			assert.Equal(t, assert.AnError, err)
		})
	})

	t.Run("Given a subscription that unsubscribes cleanly", func(t *testing.T) {
		t.Parallel()

		t.Run("Should stop tracking it once its watcher stops", func(t *testing.T) {
			t.Parallel()

			suite := WhenSubscribingToMintProposalCreatedEventsBeforeEach(t)

			subscription := newStubSubscription()
			suite.subscriptionProvider.EXPECT().
				StartMintProposalCreatedSubscription(gomock.Any(), testChainId, collectionAddress).
				Return(make(chan *abi.AgentCollectionV1MintProposalCreated), subscription, nil)

			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionOpened(gomock.Any(), mintProposalCreatedEventName, testChainId)

			stoppedTracking := make(chan struct{})
			suite.chainEventMetrics.EXPECT().
				RecordSubscriptionClosed(gomock.Any(), mintProposalCreatedEventName, testChainId).
				Do(func(context.Context, string, uint64) { close(stoppedTracking) })

			err := suite.sut.Subscribe(context.Background(), testChainId, collectionAddress)
			assert.NoError(t, err)

			// A nil error is a clean unsubscribe, the path that previously left the subscription tracked forever.
			subscription.errors <- nil

			select {
			case <-stoppedTracking:
			case <-time.After(time.Second):
				assert.Fail(t, "the subscription was still tracked after its watcher stopped")
			}
		})
	})
}
