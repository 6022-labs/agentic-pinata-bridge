package services_test

import (
	"context"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_ipfs_check_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGettingHostNodeIdsTestingSuite struct {
	sut    *services.IpfsCheckRequester
	client *clients_mocks.MockIpfsCheckClientInterface
}

func WhenGettingHostNodeIdsBeforeEach(t *testing.T) *WhenGettingHostNodeIdsTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockIpfsCheckClientInterface(mockController)

	return &WhenGettingHostNodeIdsTestingSuite{
		sut:    services.NewIpfsCheckRequester(client),
		client: client,
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
}
