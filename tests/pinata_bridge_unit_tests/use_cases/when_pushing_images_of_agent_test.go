package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPushingMissingImagesOfAgentTestingSuite struct {
	sut *use_cases.PushMissingImagesOfAgent

	agentImagesPinner *interfaces_mocks.MockAgentImagesPinnerInterface
}

func WhenPushingMissingImagesOfAgentBeforeEach(t *testing.T) *WhenPushingMissingImagesOfAgentTestingSuite {
	mockController := gomock.NewController(t)
	agentImagesPinner := interfaces_mocks.NewMockAgentImagesPinnerInterface(mockController)

	return &WhenPushingMissingImagesOfAgentTestingSuite{
		sut:               use_cases.NewPushMissingImagesOfAgent(agentImagesPinner),
		agentImagesPinner: agentImagesPinner,
	}
}

func TestWhenPushingMissingImagesOfAgent(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid request", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImagesOfAgentTestingSuite) {
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), testChainId, common.HexToAddress(testCollectionAddress), *big.NewInt(123)).
				Return(nil)
		}

		t.Run("Should hand the parsed values to the agentImagesPinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			initSuite(suite)

			response, err := suite.sut.Execute(context.Background(), &requests.PushMissingImagesOfAgentRequest{
				CollectionRequest: requests.CollectionRequest{
					ChainId:                testChainIdString,
					AgentCollectionAddress: testCollectionAddress,
				},
				AgentCollectionTokenId: "123",
			})

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	})

	t.Run("Given the agentImagesPinner fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImagesOfAgentTestingSuite) {
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(apperrors.NewUnavailableError("agent_images_read_failed", "upstream request failed"))
		}

		t.Run("Should return its error untouched", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background(), &requests.PushMissingImagesOfAgentRequest{
				CollectionRequest: requests.CollectionRequest{
					ChainId:                testChainIdString,
					AgentCollectionAddress: testCollectionAddress,
				},
				AgentCollectionTokenId: "123",
			})

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "agent_images_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given an empty request", func(t *testing.T) {
		t.Parallel()

		t.Run("Should reject the request before touching the agentImagesPinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			_, err := suite.sut.Execute(context.Background(), &requests.PushMissingImagesOfAgentRequest{})

			var validationError *apperrors.ValidationError
			assert.ErrorAs(t, err, &validationError)
		})
	})
}
