package use_cases_test

import (
	"math/big"
	"testing"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022protocol/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/services_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPushingAgentImageFromAgentTokenIdTestSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                   *services_mocks.MockPinataRequesterInterface
	agenticAIAgentCollectionRequester *services_mocks.MockAgenticAIAgentCollectionRequesterInterface
}

func WhenPushingAgentImageFromAgentTokenIdBeforeEach(t *testing.T) *WhenPushingAgentImageFromAgentTokenIdTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := services_mocks.NewMockPinataRequesterInterface(mockController)
	agenticAIAgentCollectionRequester := services_mocks.NewMockAgenticAIAgentCollectionRequesterInterface(mockController)

	sut := use_cases.NewPushAgentImageCidToPinata(
		pinataRequester,
		agenticAIAgentCollectionRequester,
	)

	return &WhenPushingAgentImageFromAgentTokenIdTestSuite{
		sut: sut,

		pinataRequester:                   pinataRequester,
		agenticAIAgentCollectionRequester: agenticAIAgentCollectionRequester,
	}
}

func TestWhenPushingAgentImageFromAgentTokenId(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAgentImageFromAgentTokenIdBeforeEach(t)
			suite.agenticAIAgentCollectionRequester.EXPECT().GetAgentImage(gomock.Any()).Return(nil, assert.AnError)

			err := suite.sut.PushFromAgentTokenId(*big.NewInt(123))
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAgentImageFromAgentTokenIdBeforeEach(t)

			testCid := "test-cid"
			suite.agenticAIAgentCollectionRequester.EXPECT().GetAgentImage(gomock.Any()).Return(&testCid, nil)
			suite.pinataRequester.EXPECT().PinCidToPinata(gomock.Any()).Return(assert.AnError)

			tokenId := big.NewInt(123)
			err := suite.sut.PushFromAgentTokenId(*tokenId)
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAgentImageFromAgentTokenIdBeforeEach(t)

			testCid := "test-cid"
			suite.agenticAIAgentCollectionRequester.EXPECT().GetAgentImage(gomock.Any()).Return(&testCid, nil)
			suite.pinataRequester.EXPECT().PinCidToPinata(gomock.Any()).Return(nil)

			err := suite.sut.PushFromAgentTokenId(*big.NewInt(123))
			assert.Equal(t, err, nil)
		})
	})
}
