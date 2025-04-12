package use_cases_test

import (
	"testing"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022protocol/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/services_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPushingAgentImageFromAgentImageCidTestSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                   *services_mocks.MockPinataRequesterInterface
	agenticAIAgentCollectionRequester *services_mocks.MockAgenticAIAgentCollectionRequesterInterface
}

func WhenPushingAgentImageFromAgentImageCidBeforeEach(t *testing.T) *WhenPushingAgentImageFromAgentImageCidTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := services_mocks.NewMockPinataRequesterInterface(mockController)
	agenticAIAgentCollectionRequester := services_mocks.NewMockAgenticAIAgentCollectionRequesterInterface(mockController)

	sut := use_cases.NewPushAgentImageCidToPinata(
		pinataRequester,
		agenticAIAgentCollectionRequester,
	)
	return &WhenPushingAgentImageFromAgentImageCidTestSuite{
		sut: sut,

		pinataRequester:                   pinataRequester,
		agenticAIAgentCollectionRequester: agenticAIAgentCollectionRequester,
	}
}

func TestWhenPushingAgentImageFromAgentImageCid(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAgentImageFromAgentImageCidBeforeEach(t)
			suite.pinataRequester.EXPECT().PinCidToPinata(gomock.Any()).Return(assert.AnError)

			err := suite.sut.PushFromAgentImageCid("test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAgentImageFromAgentImageCidBeforeEach(t)
			suite.pinataRequester.EXPECT().PinCidToPinata(gomock.Any()).Return(nil)

			err := suite.sut.PushFromAgentImageCid("test-cid")
			assert.Equal(t, err, nil)
		})
	})
}
