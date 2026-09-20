package pinata_bridge_mvc_unit_tests_test

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWhenPushingImagesOfMintProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid chain, collection and proposal id", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImagesOfAgentTestingSuite) {
			suite.mintProposalImagesPinner.EXPECT().
				Pin(gomock.Any(), uint64(80002), common.HexToAddress(validCollectionAddress), *big.NewInt(7)).
				Return(nil)
		}

		t.Run("Should reach the use case and return 204", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			initSuite(suite)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_images_of_mint_proposal/80002/"+validCollectionAddress+"/7",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)
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
