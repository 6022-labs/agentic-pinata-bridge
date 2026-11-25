package services_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"testing/synctest"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

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
			suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError).
				Times(1)

			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeEmpty, int64(1))
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

	t.Run("Given the host lookup is unavailable then finds an address", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			gomock.InOrder(
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, "")),
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return([]string{testHostNodeId}, nil),
			)
			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(2))
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
		}

		t.Run("Should retry the lookup and pin with the host address", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningACidBeforeEach(t)
				initSuite(suite)

				err := suite.sut.Pin(context.Background(), "test-cid")

				assert.NoError(t, err)
			})
		})
	})

	t.Run("Given the host lookup stays unavailable", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			gomock.InOrder(
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, "")),
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, "")),
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, "")),
			)
			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(3))
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, false, gomock.Any())
		}

		t.Run("Should pin without host addresses after three attempts", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningACidBeforeEach(t)
				initSuite(suite)

				err := suite.sut.Pin(context.Background(), "test-cid")

				assert.NoError(t, err)
			})
		})
	})

	t.Run("Given the host lookup times out", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			gomock.InOrder(
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, context.DeadlineExceeded),
			)
			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(1))
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, false, gomock.Any())
		}

		t.Run("Should pin without host addresses and not retry the lookup", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningACidBeforeEach(t)
				initSuite(suite)

				err := suite.sut.Pin(context.Background(), "test-cid")

				assert.NoError(t, err)
			})
		})
	})

	t.Run("Given the host lookup rejects the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningACidTestSuite) {
			gomock.InOrder(
				suite.ipfsCheckRequester.EXPECT().GetHostNodeIds(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("not immutable")),
			)
			suite.pinMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(1))
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			suite.pinMetrics.EXPECT().RecordPin(
				gomock.Any(), metrics_interfaces.PinOutcomePinned, false, gomock.Any())
		}

		t.Run("Should pin without host addresses and not retry the lookup", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningACidBeforeEach(t)
				initSuite(suite)

				err := suite.sut.Pin(context.Background(), "test-cid")

				assert.NoError(t, err)
			})
		})
	})
}
