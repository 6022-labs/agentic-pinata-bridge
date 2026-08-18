package services_test

import (
	"context"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingImageFromCidTestSuite struct {
	sut *services.PushAgentImageCidToPinata

	pinataRequester                  *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	pinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingImageFromCidBeforeEach(t *testing.T) *WhenPushingImageFromCidTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController)

	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	sut := services.NewPushAgentImageCidToPinata(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		pinataRequester,
		ipfsCheckRequester,
		agentCollectionRequester,
		agentCollectionsManagerRequester,
		pinMetrics,
	)
	return &WhenPushingImageFromCidTestSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pinMetrics:                       pinMetrics,
	}
}

func TestWhenPushingImageFromCid(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageFromCidBeforeEach(t)
			// First call returns error, second call returns error as well
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError).Times(2)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			// The retry drops the host addresses, so the second pin reports withHostAddresses=false.
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())

			err := suite.sut.PushFromCid(context.Background(), "test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageFromCidBeforeEach(t)
			// Only one call, returns nil
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())

			err := suite.sut.PushFromCid(context.Background(), "test-cid")
			assert.Equal(t, err, nil)
		})
	})

	t.Run("Given the host lookup found no addresses", func(t *testing.T) {
		t.Parallel()

		t.Run("Should pin once and not retry the identical request", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageFromCidBeforeEach(t)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError).Times(1)

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeEmpty, int64(3))
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())

			err := suite.sut.PushFromCid(context.Background(), "test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})
}
