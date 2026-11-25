package services_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"testing/synctest"
	"time"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_ipfs_check_mocks/clients_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_ipfs_check_mocks/metrics_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenGettingHostNodeIdsTestingSuite struct {
	sut               *services.IpfsCheckRequester
	client            *clients_mocks.MockIpfsCheckClientInterface
	hostLookupMetrics *interfaces_mocks.MockHostLookupMetricsInterface
}

func WhenGettingHostNodeIdsBeforeEach(t *testing.T) *WhenGettingHostNodeIdsTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockIpfsCheckClientInterface(mockController)
	hostLookupMetrics := interfaces_mocks.NewMockHostLookupMetricsInterface(mockController)

	return &WhenGettingHostNodeIdsTestingSuite{
		sut:               services.NewIpfsCheckRequester(zap.NewNop(), client, hostLookupMetrics),
		client:            client,
		hostLookupMetrics: hostLookupMetrics,
	}
}

func TestWhenGettingHostNodeIds(t *testing.T) {
	t.Parallel()

	t.Run("Given several reachable providers", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
				{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				{ID: "peer-2", ConnectionMaddrs: []string{"/ip4/2.2.2.2/tcp/4001", "/ip4/3.3.3.3/tcp/4001"}},
			}, nil)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
		}

		t.Run("Should return their peer ids, not their multiaddrs", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingHostNodeIdsBeforeEach(t)
			initSuite(suite)

			hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.Equal(t, []string{"peer-1", "peer-2"}, hostNodeIds)
		})
	})

	t.Run("Given a provider was never reached", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
				{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				{ID: "peer-2", ConnectionError: "dial refused"},
				{ID: "", ConnectionMaddrs: []string{"/ip4/4.4.4.4/tcp/4001"}},
			}, nil)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
		}

		t.Run("Should skip it, since pinata gains nothing from an unreachable node", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingHostNodeIdsBeforeEach(t)
			initSuite(suite)

			hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.Equal(t, []string{"peer-1"}, hostNodeIds)
		})
	})

	t.Run("Given the same provider is reported by several sources", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
				{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}, Source: "dht"},
				{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}, Source: "bitswap"},
			}, nil)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
		}

		t.Run("Should report its peer id once", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingHostNodeIdsBeforeEach(t)
			initSuite(suite)

			hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.Equal(t, []string{"peer-1"}, hostNodeIds)
		})
	})

	t.Run("Given no providers are returned", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{}, nil)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeEmpty, int64(1))
		}

		t.Run("Should return an empty, non-nil slice", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingHostNodeIdsBeforeEach(t)
			initSuite(suite)

			hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, hostNodeIds)
			assert.Empty(t, hostNodeIds)
		})
	})

	t.Run("Given the client fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return(nil, assert.AnError)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(1))
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingHostNodeIdsBeforeEach(t)
			initSuite(suite)

			hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, hostNodeIds)
		})
	})

	t.Run("Given ipfs-check is unavailable then finds a provider", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().
					Check(gomock.Any(), "QmHash").
					Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, 0, "")),
				suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
					{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				}, nil),
			)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(2))
		}

		t.Run("Should return the peer id found on the second attempt", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.NoError(t, err)
				assert.Equal(t, []string{"peer-1"}, hostNodeIds)
			})
		})
	})

	t.Run("Given the connection drops then ipfs-check finds a provider", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return(nil, &url.Error{Op: "Post", Err: io.EOF}),
				suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
					{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				}, nil),
			)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(2))
		}

		t.Run("Should return the peer id found on the second attempt", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.NoError(t, err)
				assert.Equal(t, []string{"peer-1"}, hostNodeIds)
			})
		})
	})

	t.Run("Given the response is truncated then ipfs-check finds a provider", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return(nil, io.ErrUnexpectedEOF),
				suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
					{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				}, nil),
			)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(2))
		}

		t.Run("Should return the peer id found on the second attempt", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.NoError(t, err)
				assert.Equal(t, []string{"peer-1"}, hostNodeIds)
			})
		})
	})

	t.Run("Given ipfs-check stays unavailable", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").
				Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusServiceUnavailable, 0, "")).
				Times(3)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(3))
		}

		t.Run("Should return the error after three attempts", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.Error(t, err)
				assert.Nil(t, hostNodeIds)
				assert.InDelta(t, 3*time.Second, time.Since(start), float64(300*time.Millisecond))
			})
		})
	})

	t.Run("Given ipfs-check times out", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").
				Return(nil, &url.Error{Op: "Post", Err: context.DeadlineExceeded}).
				Times(1)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(1))
		}

		t.Run("Should return the error without retrying", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.Error(t, err)
				assert.Nil(t, hostNodeIds)
			})
		})
	})

	t.Run("Given ipfs-check refuses the request", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingHostNodeIdsTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").
				Return(nil, apperrors.NewUpstreamStatusError("ipfs-check", http.StatusBadRequest, 0, "")).
				Times(1)
			suite.hostLookupMetrics.EXPECT().RecordHostLookup(
				gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, int64(1))
		}

		t.Run("Should return the error without retrying", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenGettingHostNodeIdsBeforeEach(t)
				initSuite(suite)

				hostNodeIds, err := suite.sut.GetHostNodeIds(context.Background(), "QmHash")

				assert.Error(t, err)
				assert.Nil(t, hostNodeIds)
			})
		})
	})
}
