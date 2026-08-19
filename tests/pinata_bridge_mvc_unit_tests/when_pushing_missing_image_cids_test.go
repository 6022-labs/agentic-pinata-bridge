package pinata_bridge_mvc_unit_tests_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWhenPushingMissingImageCids(t *testing.T) {
	t.Parallel()

	t.Run("Given no chain is configured", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImagesOfAgentTestingSuite) {
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), gomock.Any(), gomock.Any(), false)
		}

		t.Run("Should sweep nothing and answer an empty 200", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			initSuite(suite)

			req := httptest.NewRequest(http.MethodPost, "/push_missing_image_cids", nil)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			body, readErr := io.ReadAll(resp.Body)
			assert.NoError(t, readErr)
			assert.Empty(t, body)
		})
	})
}
