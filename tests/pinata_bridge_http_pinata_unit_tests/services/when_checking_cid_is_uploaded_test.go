package services_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_pinata_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenCheckingCidIsUploadedTestingSuite struct {
	sut    *services.PinataRequester
	client *clients_mocks.MockClientWithResponsesInterface
}

func WhenCheckingCidIsUploadedBeforeEach(t *testing.T) *WhenCheckingCidIsUploadedTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockClientWithResponsesInterface(mockController)

	return &WhenCheckingCidIsUploadedTestingSuite{
		sut:    services.NewPinataRequester(client, &settings.PinataSettings{RequestTimeout: testRequestTimeout}),
		client: client,
	}
}

// listFilesResponseWith fills the generated anonymous JSON200 struct the only sane way: through the wire shape.
func listFilesResponseWith(t *testing.T, fileCount int) *clients.ListFilesResponse {
	t.Helper()

	files := make([]json.RawMessage, fileCount)
	for i := range files {
		files[i] = json.RawMessage(`{}`)
	}

	body, err := json.Marshal(map[string]any{"data": map[string]any{"files": files}})
	assert.NoError(t, err)

	response := &clients.ListFilesResponse{
		Body:         body,
		HTTPResponse: &http.Response{StatusCode: http.StatusOK},
	}
	assert.NoError(t, json.Unmarshal(body, &response.JSON200))

	return response
}

func TestWhenCheckingCidIsUploaded(t *testing.T) {
	t.Parallel()

	t.Run("Given exactly one file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					_ context.Context,
					network clients.ListFilesParamsNetwork,
					params *clients.ListFilesParams,
					_ ...clients.RequestEditorFn,
				) (*clients.ListFilesResponse, error) {
					assert.Equal(t, clients.Public, network)
					assert.NotNil(t, params.Cid)
					assert.Equal(t, testCid, *params.Cid)

					return listFilesResponseWith(t, 1), nil
				},
			)
		}

		t.Run("Should report the cid as uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.True(t, *uploaded)
		})
	})

	t.Run("Given no file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(listFilesResponseWith(t, 0), nil)
		}

		t.Run("Should report the cid as not uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.False(t, *uploaded)
		})
	})

	t.Run("Given more than one file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(listFilesResponseWith(t, 2), nil)
		}

		t.Run("Should report the cid as not uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.False(t, *uploaded)
		})
	})

	t.Run("Given pinata answers a non-2xx status", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(&clients.ListFilesResponse{
					Body:         []byte(`{"error":"unauthorized"}`),
					HTTPResponse: &http.Response{StatusCode: http.StatusUnauthorized},
				}, nil)
		}

		t.Run("Should return an error carrying the status and body", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.Nil(t, uploaded)
			assert.ErrorContains(t, err, "401")
			assert.ErrorContains(t, err, "unauthorized")
		})
	})

	t.Run("Given the transport fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, uploaded)
		})
	})

	t.Run("Given pinata never answers", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					ctx context.Context,
					_ clients.ListFilesParamsNetwork,
					_ *clients.ListFilesParams,
					_ ...clients.RequestEditorFn,
				) (*clients.ListFilesResponse, error) {
					<-ctx.Done()

					return nil, ctx.Err()
				},
			)
		}

		t.Run("Should give up after the request timeout", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			isUploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

			assert.Nil(t, isUploaded)
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		})
	})

	t.Run("Given pinata is rate limiting then answers", func(t *testing.T) {
		t.Parallel()

		initSuite := func(t *testing.T, suite *WhenCheckingCidIsUploadedTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&clients.ListFilesResponse{
						HTTPResponse: &http.Response{StatusCode: http.StatusTooManyRequests},
					}, nil),
				suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(listFilesResponseWith(t, 1), nil),
			)
		}

		t.Run("Should retry and report the cid as uploaded", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenCheckingCidIsUploadedBeforeEach(t)
				initSuite(t, suite)

				isUploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

				if assert.NoError(t, err) {
					assert.True(t, *isUploaded)
				}
			})
		})
	})

	t.Run("Given pinata fails with a server error", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().ListFilesWithResponse(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(&clients.ListFilesResponse{
					HTTPResponse: &http.Response{StatusCode: http.StatusServiceUnavailable},
				}, nil).
				Times(1)
		}

		t.Run("Should return the error without retrying", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenCheckingCidIsUploadedBeforeEach(t)
				initSuite(suite)

				isUploaded, err := suite.sut.IsCidUploaded(context.Background(), testCid)

				assert.Nil(t, isUploaded)
				assert.ErrorContains(t, err, "503")
			})
		})
	})
}
