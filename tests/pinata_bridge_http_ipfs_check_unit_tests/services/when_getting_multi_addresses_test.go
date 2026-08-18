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

type WhenGettingMultiAddressesTestingSuite struct {
	sut    *services.IpfsCheckRequester
	client *clients_mocks.MockIpfsCheckClientInterface
}

func WhenGettingMultiAddressesBeforeEach(t *testing.T) *WhenGettingMultiAddressesTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockIpfsCheckClientInterface(mockController)

	return &WhenGettingMultiAddressesTestingSuite{
		sut:    services.NewIpfsCheckRequester(client),
		client: client,
	}
}

func TestWhenGettingMultiAddresses(t *testing.T) {
	t.Parallel()

	t.Run("Given several providers each expose multiaddrs", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingMultiAddressesTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{
				{ID: "peer-1", ConnectionMaddrs: []string{"/ip4/1.1.1.1/tcp/4001"}},
				{ID: "peer-2", ConnectionMaddrs: []string{"/ip4/2.2.2.2/tcp/4001", "/ip4/3.3.3.3/tcp/4001"}},
			}, nil)
		}

		t.Run("Should flatten every provider's multiaddrs", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingMultiAddressesBeforeEach(t)
			initSuite(suite)

			addresses, err := suite.sut.GetMultiAddresses(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.Equal(t, []string{
				"/ip4/1.1.1.1/tcp/4001",
				"/ip4/2.2.2.2/tcp/4001",
				"/ip4/3.3.3.3/tcp/4001",
			}, addresses)
		})
	})

	t.Run("Given no providers are returned", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingMultiAddressesTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return([]models.ExternalCheckResponse{}, nil)
		}

		t.Run("Should return an empty, non-nil slice", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingMultiAddressesBeforeEach(t)
			initSuite(suite)

			addresses, err := suite.sut.GetMultiAddresses(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, addresses)
			assert.Empty(t, addresses)
		})
	})

	t.Run("Given the client fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenGettingMultiAddressesTestingSuite) {
			suite.client.EXPECT().Check(gomock.Any(), "QmHash").Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingMultiAddressesBeforeEach(t)
			initSuite(suite)

			addresses, err := suite.sut.GetMultiAddresses(context.Background(), "QmHash")

			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, addresses)
		})
	})
}
