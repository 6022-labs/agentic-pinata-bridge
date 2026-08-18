package pinata_bridge_mvc_unit_tests_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWhenPushingImagesOfMintProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid chain, collection and proposal id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should reach the use case and return 200", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			suite.agentCollectionRequester.EXPECT().
				GetMintProposalImages(gomock.Any(), uint64(80002), gomock.Any(), gomock.Any()).
				Return(nil, nil)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), gomock.Any(), gomock.Any(), false)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_images_of_mint_proposal/80002/"+validCollectionAddress+"/7",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})
	})

	t.Run("Given a proposal id that is not a number", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return 400 naming the mintProposalId field", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_images_of_mint_proposal/80002/"+validCollectionAddress+"/not-a-number",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			var body map[string]string
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
			assert.Equal(t, "mintProposalId", body["field"])
		})
	})
}
