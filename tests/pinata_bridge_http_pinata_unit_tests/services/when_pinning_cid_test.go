package services_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_pinata_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPinningCidTestingSuite struct {
	sut    *services.PinataRequester
	client *clients_mocks.MockClientWithResponsesInterface
}

func WhenPinningCidBeforeEach(t *testing.T) *WhenPinningCidTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockClientWithResponsesInterface(mockController)

	return &WhenPinningCidTestingSuite{
		sut:    services.NewPinataRequester(client),
		client: client,
	}
}

func pinByCidResponseWith(statusCode int, body string) *clients.PinByCidResponse {
	return &clients.PinByCidResponse{
		Body:         []byte(body),
		HTTPResponse: &http.Response{StatusCode: statusCode},
	}
}

func TestWhenPinningCid(t *testing.T) {
	t.Parallel()

	t.Run("Given host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					_ context.Context,
					body clients.PinByCidJSONRequestBody,
					_ ...clients.RequestEditorFn,
				) (*clients.PinByCidResponse, error) {
					assert.Equal(t, testCid, body.Cid)
					assert.NotNil(t, body.HostNodes)
					assert.Equal(t, []string{testHostNode}, *body.HostNodes)

					return pinByCidResponseWith(http.StatusOK, `{}`), nil
				},
			)
		}

		t.Run("Should pin the cid carrying the host nodes", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

			assert.NoError(t, err)
		})
	})

	t.Run("Given no host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					_ context.Context,
					body clients.PinByCidJSONRequestBody,
					_ ...clients.RequestEditorFn,
				) (*clients.PinByCidResponse, error) {
					assert.Equal(t, testCid, body.Cid)
					assert.Nil(t, body.HostNodes)

					return pinByCidResponseWith(http.StatusOK, `{}`), nil
				},
			)
		}

		t.Run("Should pin the cid without host nodes", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.NoError(t, err)
		})
	})

	t.Run("Given pinata rejects the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).
				Return(pinByCidResponseWith(http.StatusBadRequest, `{"error":"Invalid IPFS hash"}`), nil)
		}

		t.Run("Should return an error carrying the status and body", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.ErrorContains(t, err, "400")
			assert.ErrorContains(t, err, "Invalid IPFS hash")
		})
	})

	t.Run("Given the transport fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.ErrorIs(t, err, assert.AnError)
		})
	})
}
