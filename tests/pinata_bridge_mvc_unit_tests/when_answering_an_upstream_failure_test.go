package pinata_bridge_mvc_unit_tests_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_mvc"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenAnsweringAnUpstreamFailureTestingSuite struct {
	app *fiber.App

	mockAgentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	mockPinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenAnsweringAnUpstreamFailureBeforeEach(t *testing.T) *WhenAnsweringAnUpstreamFailureTestingSuite {
	mockController := gomock.NewController(t)

	mockPinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	mockIpfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	mockAgentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	mockAgentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
		mockController,
	)
	mockPinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)
	mockPinTracer := newNoopPinTracer(mockController)

	cidPinner := services.NewCidPinner(
		zap.NewNop(),
		mockPinataRequester,
		mockIpfsCheckRequester,
		mockPinMetrics,
		mockPinTracer,
	)
	pushMissingImagesOfAgent := use_cases.NewPushMissingImagesOfAgent(
		zap.NewNop(), cidPinner, mockAgentCollectionRequester, mockPinataRequester, mockPinMetrics,
	)

	controller := pinata_bridge_mvc.NewPinataPushController(
		use_cases.NewPushMissingImageCids(
			zap.NewNop(), mockAgentCollectionRequester, mockPinMetrics,
			// One configured chain, so the sweep actually reaches the manager requester.
			settings.NewChainsSettingsFromChainIds([]uint64{80002}),
			mockAgentCollectionsManagerRequester, pushMissingImagesOfAgent, mockPinTracer,
		),
		pushMissingImagesOfAgent,
		use_cases.NewPushImagesOfMintProposal(zap.NewNop(), cidPinner, mockAgentCollectionRequester, mockPinMetrics),
	)

	app := fiber.New()
	controller.RegisterRoutes(app)

	return &WhenAnsweringAnUpstreamFailureTestingSuite{
		app:                                  app,
		mockAgentCollectionsManagerRequester: mockAgentCollectionsManagerRequester,
		mockPinMetrics:                       mockPinMetrics,
	}
}

func TestWhenAnsweringAnUpstreamFailure(t *testing.T) {
	t.Parallel()

	t.Run("Given the chain read fails with an error quoting the rpc url", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenAnsweringAnUpstreamFailureTestingSuite) {
			suite.mockAgentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), uint64(80002)).
				Return(nil, errors.New(`Post "https://secret-key@rpc.example": dial tcp: i/o timeout`))
			suite.mockPinMetrics.EXPECT().RecordSweep(gomock.Any(), gomock.Any(), gomock.Any(), true)
		}

		t.Run("Should answer 502 carrying the code but never the upstream text", func(t *testing.T) {
			t.Parallel()

			suite := WhenAnsweringAnUpstreamFailureBeforeEach(t)
			initSuite(suite)

			req := httptest.NewRequest(http.MethodPost, "/push_missing_image_cids", nil)
			resp, err := suite.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadGateway, resp.StatusCode)

			body, readErr := io.ReadAll(resp.Body)
			assert.NoError(t, readErr)
			assert.Contains(t, string(body), "collections_read_failed")

			// These routes are unauthenticated and an rpc url can carry an api key.
			assert.NotContains(t, string(body), "secret-key")
			assert.NotContains(t, string(body), "rpc.example")
		})
	})
}
