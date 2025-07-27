package use_cases_test

import (
	"testing"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022protocol/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/services_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingImageFromCidTestSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                  *services_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *services_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *services_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *services_mocks.MockAgentCollectionsManagerRequesterInterface
}

func WhenPushingImageFromCidBeforeEach(t *testing.T) *WhenPushingImageFromCidTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := services_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := services_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := services_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := services_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController)

	sut := use_cases.NewPushAgentImageCidToPinata(
		zap.NewNop(),
		pinataRequester,
		ipfsCheckRequester,
		agentCollectionRequester,
		agentCollectionsManagerRequester,
	)
	return &WhenPushingImageFromCidTestSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
	}
}

func TestWhenPushingImageFromCid(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageFromCidBeforeEach(t)
			// First call returns error, second call returns error as well
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(assert.AnError).Times(2)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			err := suite.sut.PushFromCid("test-cid")
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageFromCidBeforeEach(t)
			// Only one call, returns nil
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			err := suite.sut.PushFromCid("test-cid")
			assert.Equal(t, err, nil)
		})
	})
}
