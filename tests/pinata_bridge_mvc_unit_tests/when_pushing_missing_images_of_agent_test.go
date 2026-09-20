package pinata_bridge_mvc_unit_tests_test

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

const validCollectionAddress = "0x1234567890123456789012345678901234567890"

type WhenPushingMissingImagesOfAgentTestingSuite struct {
	app *fiber.App

	agentImagesPinner        *interfaces_mocks.MockAgentImagesPinnerInterface
	mintProposalImagesPinner *interfaces_mocks.MockMintProposalImagesPinnerInterface
	pinMetrics               *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingMissingImagesOfAgentBeforeEach(t *testing.T) *WhenPushingMissingImagesOfAgentTestingSuite {
	mockController := gomock.NewController(t)

	agentImagesPinner := interfaces_mocks.NewMockAgentImagesPinnerInterface(mockController)
	mintProposalImagesPinner := interfaces_mocks.NewMockMintProposalImagesPinnerInterface(mockController)
	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	controller := pinata_bridge_mvc.NewPinataPushController(
		use_cases.NewPushMissingImageCids(
			zap.NewNop(),
			interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController),
			pinMetrics,
			settings.NewChainsSettingsFromChainIds(nil),
			interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController),
			agentImagesPinner,
			newNoopPinTracer(mockController),
		),
		use_cases.NewPushMissingImagesOfAgent(agentImagesPinner),
		use_cases.NewPushImagesOfMintProposal(mintProposalImagesPinner),
	)

	app := fiber.New()
	controller.RegisterRoutes(app)

	return &WhenPushingMissingImagesOfAgentTestingSuite{
		app:                      app,
		agentImagesPinner:        agentImagesPinner,
		mintProposalImagesPinner: mintProposalImagesPinner,
		pinMetrics:               pinMetrics,
	}
}

func TestWhenPushingMissingImagesOfAgent(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid chain, collection and token id", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImagesOfAgentTestingSuite) {
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), uint64(80002), common.HexToAddress(validCollectionAddress), *big.NewInt(123)).
				Return(nil)
		}

		t.Run("Should reach the use case and return 204", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			initSuite(suite)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_missing_images_of_agent/80002/"+validCollectionAddress+"/123",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			body, readErr := io.ReadAll(resp.Body)
			assert.NoError(t, readErr)
			assert.Empty(t, body)
		})
	})

	t.Run("Given a non-numeric chain id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return 400 without reaching the use case", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_missing_images_of_agent/not-a-chain/"+validCollectionAddress+"/123",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			var body map[string]string
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
			assert.Equal(t, "chainId", body["field"])
			assert.Equal(t, "chainId is invalid", body["message"])
		})
	})

	t.Run("Given a collection address that is not hex", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return 400 with the agentCollectionAddress field", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_missing_images_of_agent/80002/not-an-address/123",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			var body map[string]string
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
			assert.Equal(t, "agentCollectionAddress", body["field"])
		})
	})

	t.Run("Given a token id that is not a number", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return 400 with the agentCollectionTokenId field", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			req := httptest.NewRequest(
				http.MethodPost,
				"/push_missing_images_of_agent/80002/"+validCollectionAddress+"/not-a-number",
				nil,
			)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			var body map[string]string
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
			assert.Equal(t, "agentCollectionTokenId", body["field"])
		})
	})
}
