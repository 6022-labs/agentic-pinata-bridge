package pinata_bridge_listeners_unit_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_listeners_mocks/metrics_mocks/interfaces_mocks"
	metrics_mocks_pin "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	interfaces_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

const otherChainId uint64 = 137

// A replaced subscription must survive until its own chain proves the replacement is live.
func TestWhenRetiringAReplacedSubscription(t *testing.T) {
	t.Parallel()

	collection := common.HexToAddress("0x1111111111111111111111111111111111111111")
	newCollection := common.HexToAddress("0x2222222222222222222222222222222222222222")

	t.Run("Given two chains and a rebuild on one of them", func(t *testing.T) {
		t.Parallel()

		t.Run("Should retire only this chain's parked subscription", func(t *testing.T) {
			t.Parallel()

			mockController := gomock.NewController(t)

			agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
				mockController,
			)
			subscriptionProvider := interfaces_mocks.NewMockMintedSubscriptionProviderInterface(mockController)
			chainEventMetrics := metrics_mocks.NewMockChainEventMetricsInterface(mockController)
			pinMetrics := metrics_mocks_pin.NewMockPinMetricsInterface(mockController)

			agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
			agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, nil).AnyTimes()

			handleMintedEvent := use_cases.NewHandleMintedEvent(use_cases.NewPushMissingImagesOfAgent(
				zap.NewNop(),
				interfaces_mocks.NewMockCidPinnerInterface(mockController),
				agentCollectionRequester,
				interfaces_mocks.NewMockPinataRequesterInterface(mockController),
				pinMetrics,
			))

			sut := pinata_bridge_listeners.NewAgentCollectionMintedListener(
				zap.NewNop(),
				settings.NewChainsSettingsFromChainIds([]uint64{testChainId, otherChainId}),
				use_cases.NewListCollectionAddresses(agentCollectionsManagerRequester),
				chainEventMetrics,
				subscriptionProvider,
				handleMintedEvent,
			)

			agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), gomock.Any()).
				Return([]common.Address{collection}, nil).Times(2)

			// Both chains rebuild, so each parks its first subscription. The other chain's parked
			// subscription is the one that must survive an event on this chain.
			otherChainParkedSubscription := newStubSubscription()
			testChainEvents := make(chan *abi.AgentCollectionV1Minted, 1)

			subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), otherChainId, []common.Address{collection}).
				Return(make(chan *abi.AgentCollectionV1Minted), otherChainParkedSubscription, nil)
			subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), testChainId, []common.Address{collection}).
				Return(make(chan *abi.AgentCollectionV1Minted), newStubSubscription(), nil)

			subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), otherChainId, []common.Address{collection, newCollection}).
				Return(make(chan *abi.AgentCollectionV1Minted), newStubSubscription(), nil)
			subscriptionProvider.EXPECT().
				StartMintedSubscription(gomock.Any(), testChainId, []common.Address{collection, newCollection}).
				Return(testChainEvents, newStubSubscription(), nil)

			chainEventMetrics.EXPECT().RecordSubscriptionOpened(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
			chainEventMetrics.EXPECT().RecordSubscriptionClosed(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

			handled := make(chan struct{}, 1)
			chainEventMetrics.EXPECT().
				RecordEvent(gomock.Any(), mintedEventName, testChainId, gomock.Any(), gomock.Any()).
				Do(func(context.Context, string, uint64, string, time.Duration) { handled <- struct{}{} })
			pinMetrics.EXPECT().RecordSweep(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

			assert.NoError(t, sut.SubscribeAll(context.Background()))
			assert.NoError(t, sut.Subscribe(context.Background(), otherChainId, newCollection))
			assert.NoError(t, sut.Subscribe(context.Background(), testChainId, newCollection))

			// An event on testChainId must only retire testChainId's parked subscription.
			testChainEvents <- &abi.AgentCollectionV1Minted{TokenId: big.NewInt(1), Raw: types.Log{BlockNumber: 100}}

			select {
			case <-handled:
			case <-time.After(2 * time.Second):
				assert.Fail(t, "the event was never handled")
			}

			assert.False(
				t,
				otherChainParkedSubscription.unsubscribed(),
				"an event on one chain retired another chain's parked subscription",
			)

			var _ ethereum.Subscription = otherChainParkedSubscription
		})
	})
}
