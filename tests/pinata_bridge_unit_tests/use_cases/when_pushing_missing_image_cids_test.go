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

type WhenPushingMissingImageCidsTestingSuite struct {
	sut *use_cases.PushAgentImageCidToPinata

	pinataRequester                  *services_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *services_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *services_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *services_mocks.MockAgentCollectionsManagerRequesterInterface
}

func WhenPushingMissingImageCidsBeforeEach(t *testing.T) *WhenPushingMissingImageCidsTestingSuite {
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
	return &WhenPushingMissingImageCidsTestingSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
	}
}

func TestWhenPushingMissingImageCids(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting all collections addresses", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			suite.agentCollectionsManagerRequester.EXPECT().GetAllCollectionAddresses().Return(nil, assert.AnError)

			err := suite.sut.PushMissingImageCids()
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given error occurs while getting all token ids", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			collectionAddress := []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			}

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			suite.agentCollectionsManagerRequester.EXPECT().GetAllCollectionAddresses().Return(collectionAddress, nil)
			for _, address := range collectionAddress {
				suite.agentCollectionRequester.EXPECT().GetAllTokenIds(address).Return(nil, assert.AnError)
			}

			err := suite.sut.PushMissingImageCids()
			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs while getting all token ids", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)

			collectionAddress := []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			}

			tokenIds := []big.Int{
				*big.NewInt(1),
				*big.NewInt(2),
			}

			imageCid := "test-cid"

			suite.agentCollectionsManagerRequester.EXPECT().GetAllCollectionAddresses().Return(collectionAddress, nil)
			for _, address := range collectionAddress {
				suite.agentCollectionRequester.EXPECT().GetAllTokenIds(address).Return(tokenIds, nil)
				for _, tokenId := range tokenIds {
					suite.agentCollectionRequester.EXPECT().GetAgentImages(address, tokenId).Return([]string{imageCid}, nil)
				}
			}

			isCidUploaded := false

			suite.pinataRequester.EXPECT().IsCidUploaded(imageCid).Return(&isCidUploaded, nil).Times(len(tokenIds))
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(nil).Times(len(tokenIds))
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return(nil, nil).AnyTimes()

			err := suite.sut.PushMissingImageCids()
			assert.NoError(t, err)
		})
	})
}
