package services_test

import (
	"context"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_pinata_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPinningCidTestingSuite struct {
	sut    *services.PinataRequester
	client *clients_mocks.MockPinataClientInterface
}

func WhenPinningCidBeforeEach(t *testing.T) *WhenPinningCidTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockPinataClientInterface(mockController)

	return &WhenPinningCidTestingSuite{
		sut:    services.NewPinataRequester(client),
		client: client,
	}
}

func TestWhenPinningCid(t *testing.T) {
	t.Parallel()

	t.Run("Given host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByHash(gomock.Any(), gomock.Any()).DoAndReturn(
				func(_ context.Context, request *models.ExternalPinByHashRequest) (*models.ExternalPinByHashResponse, error) {
					assert.Equal(t, "QmHash", request.HashToPin)
					assert.NotNil(t, request.PinataOptions)
					assert.Equal(t, []string{"/ip4/127.0.0.1/tcp/4001"}, request.PinataOptions.HostNodes)
					return &models.ExternalPinByHashResponse{}, nil
				},
			)
		}

		t.Run("Should pin the cid carrying the host nodes in the options", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), "QmHash", []string{"/ip4/127.0.0.1/tcp/4001"})

			assert.NoError(t, err)
		})
	})

	t.Run("Given no host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByHash(gomock.Any(), gomock.Any()).DoAndReturn(
				func(_ context.Context, request *models.ExternalPinByHashRequest) (*models.ExternalPinByHashResponse, error) {
					assert.Equal(t, "QmHash", request.HashToPin)
					assert.Nil(t, request.PinataOptions)
					return &models.ExternalPinByHashResponse{}, nil
				},
			)
		}

		t.Run("Should pin the cid without options", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), "QmHash", nil)

			assert.NoError(t, err)
		})
	})

	t.Run("Given the client fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByHash(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), "QmHash", nil)

			assert.ErrorIs(t, err, assert.AnError)
		})
	})
}
