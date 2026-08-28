package services_test

import (
	"context"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

// testHostNodeId is a libp2p peer id, the shape pinata's host_nodes expects.
const testHostNodeId = "12D3KooWEyoppNCUx8Yx66oV9fJnriXwCcXwDDUA2kj6vnc6iDEg"

type WhenPinningACidTestSuite struct {
	sut *services.CidPinner

	pinataRequester    *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester *interfaces_mocks.MockIpfsCheckRequesterInterface
	pinMetrics         *metrics_mocks.MockPinMetricsInterface
}

func WhenPinningACidBeforeEach(t *testing.T) *WhenPinningACidTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)

	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	sut := services.NewCidPinner(
		zap.NewNop(),
		pinataRequester,
		ipfsCheckRequester,
		pinMetrics,
		newNoopPinTracer(mockController),
	)
	return &WhenPinningACidTestSuite{
		sut: sut,

		pinataRequester:    pinataRequester,
		ipfsCheckRequester: ipfsCheckRequester,
		pinMetrics:         pinMetrics,
	}
}

func TestWhenPinningACid(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			// First call returns error, second call returns error as well
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError).
				Times(2)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return([]string{testHostNodeId}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			// The retry drops the host addresses, so the second pin reports withHostAddresses=false.
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningACidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(context.Background(), "test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			// Only one call, returns nil
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return([]string{testHostNodeId}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
		}

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningACidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(context.Background(), "test-cid")
			assert.Equal(t, err, nil)
		})
	})

	t.Run("Given the host lookup found no addresses", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError).
				Times(1)

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeEmpty, int64(3))
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
		}

		t.Run("Should pin once and not retry the identical request", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningACidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(context.Background(), "test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})
}
