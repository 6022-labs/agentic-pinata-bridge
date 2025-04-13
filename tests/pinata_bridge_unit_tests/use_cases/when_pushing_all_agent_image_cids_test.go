package use_cases_test

import (
	"math/big"
	"testing"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022protocol/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/services_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingAllAgentImageCidsTestSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                   *services_mocks.MockPinataRequesterInterface
	ipfsCheckRequester                *services_mocks.MockIpfsCheckRequesterInterface
	agenticAIAgentCollectionRequester *services_mocks.MockAgenticAIAgentCollectionRequesterInterface
}

func WhenPushingAllAgentImageCidsBeforeEach(t *testing.T) *WhenPushingAllAgentImageCidsTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := services_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := services_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agenticAIAgentCollectionRequester := services_mocks.NewMockAgenticAIAgentCollectionRequesterInterface(mockController)

	sut := use_cases.NewPushAgentImageCidToPinata(
		zap.NewNop(),
		pinataRequester,
		ipfsCheckRequester,
		agenticAIAgentCollectionRequester,
	)
	return &WhenPushingAllAgentImageCidsTestSuite{
		sut: sut,

		pinataRequester:                   pinataRequester,
		ipfsCheckRequester:                ipfsCheckRequester,
		agenticAIAgentCollectionRequester: agenticAIAgentCollectionRequester,
	}
}

func TestWhenPushingAllAgentImageCids(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting all token ids", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAllAgentImageCidsBeforeEach(t)
			suite.agenticAIAgentCollectionRequester.EXPECT().GetAllTokenIds().Return(nil, assert.AnError)

			err := suite.sut.PushAllAgentImageCids()
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while getting all token ids", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingAllAgentImageCidsBeforeEach(t)

			tokenIds := []big.Int{
				*big.NewInt(1),
				*big.NewInt(2),
			}

			imageCid := "test-cid"

			suite.agenticAIAgentCollectionRequester.EXPECT().GetAllTokenIds().Return(tokenIds, nil)
			suite.agenticAIAgentCollectionRequester.EXPECT().GetAgentImage(gomock.Any()).Return(&imageCid, nil).Times(len(tokenIds))
			suite.pinataRequester.EXPECT().PinCidToPinata(gomock.Any(), gomock.Any()).Return(nil).Times(len(tokenIds))
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return(nil, nil).AnyTimes()

			err := suite.sut.PushAllAgentImageCids()
			assert.NoError(t, err)
		})
	})
}
