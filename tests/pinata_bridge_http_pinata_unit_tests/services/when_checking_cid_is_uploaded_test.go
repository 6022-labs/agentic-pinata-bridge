package services_test

import (
	"encoding/json"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_pinata_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenCheckingCidIsUploadedTestingSuite struct {
	sut    *services.PinataRequester
	client *clients_mocks.MockPinataClientInterface
}

func WhenCheckingCidIsUploadedBeforeEach(t *testing.T) *WhenCheckingCidIsUploadedTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockPinataClientInterface(mockController)

	return &WhenCheckingCidIsUploadedTestingSuite{
		sut:    services.NewPinataRequester(client),
		client: client,
	}
}

func queryFilesResponseWith(t *testing.T, fileCount int) *models.ExternalQueryFilesResponse {
	t.Helper()

	files := make([]json.RawMessage, fileCount)
	for i := range files {
		files[i] = json.RawMessage(`{}`)
	}
	payload, err := json.Marshal(map[string]any{"data": map[string]any{"files": files}})
	assert.NoError(t, err)

	var resp models.ExternalQueryFilesResponse
	assert.NoError(t, json.Unmarshal(payload, &resp))
	return &resp
}

func TestWhenCheckingCidIsUploaded(t *testing.T) {
	t.Parallel()

	t.Run("Given exactly one file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().QueryFileByCid("QmHash").Return(queryFilesResponseWith(t, 1), nil)
		}

		t.Run("Should report the cid as uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded("QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.True(t, *uploaded)
		})
	})

	t.Run("Given no file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().QueryFileByCid("QmHash").Return(queryFilesResponseWith(t, 0), nil)
		}

		t.Run("Should report the cid as not uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded("QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.False(t, *uploaded)
		})
	})

	t.Run("Given more than one file matches the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().QueryFileByCid("QmHash").Return(queryFilesResponseWith(t, 2), nil)
		}

		t.Run("Should report the cid as not uploaded", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded("QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, uploaded)
			assert.False(t, *uploaded)
		})
	})

	t.Run("Given the client fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidIsUploadedTestingSuite) {
			suite.client.EXPECT().QueryFileByCid("QmHash").Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidIsUploadedBeforeEach(t)
			initSuite(suite)

			uploaded, err := suite.sut.IsCidUploaded("QmHash")

			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, uploaded)
		})
	})
}
