package use_cases_test

import (
	"math/big"
	"testing"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022protocol/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/services_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingMissingImagesOfAgentTestingSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                  *services_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *services_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *services_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *services_mocks.MockAgentCollectionsManagerRequesterInterface
}

func WhenPushingMissingImagesOfAgentBeforeEach(t *testing.T) *WhenPushingMissingImagesOfAgentTestingSuite {
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

	return &WhenPushingMissingImagesOfAgentTestingSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
	}
}

func TestWhenPushingImagesOfAgent(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

			err := suite.sut.PushMissingImagesOfAgent(common.HexToAddress(""), *big.NewInt(123))

			assert.NotNil(t, err)
		})
	})

	t.Run("Given error occurs while checking if cid is already pinned", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(testCid).Return(nil, assert.AnError)

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(common.HexToAddress(""), *tokenId)

			assert.NotNil(t, err)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			testCid := "test-cid"
			isCidUploaded := false
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(assert.AnError).Times(2)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(common.HexToAddress(""), *tokenId)

			assert.NotNil(t, err)
		})
	})

	t.Run("Given image cid is already pinned", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			testCid := "test-cid"
			isCidUploaded := true
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(testCid).Return(&isCidUploaded, nil)

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(common.HexToAddress(""), *tokenId)

			assert.Nil(t, err)
		})
	})

	t.Run("Given image cid is not yet pinned", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			testCid := "test-cid"
			isCidUploaded := false
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			err := suite.sut.PushMissingImagesOfAgent(common.HexToAddress(""), *big.NewInt(123))

			assert.Nil(t, err)
		})
	})
}
