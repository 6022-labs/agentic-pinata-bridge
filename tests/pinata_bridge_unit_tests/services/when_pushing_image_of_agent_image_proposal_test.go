package services_test

import (
	"math/big"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingImageOfAgentImageProposalTestSuite struct {
	sut *services.PushAgentImageCidToPinata

	pinataRequester                  *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
}

func WhenPushingImageOfAgentImageProposalBeforeEach(t *testing.T) *WhenPushingImageOfAgentImageProposalTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController)

	sut := services.NewPushAgentImageCidToPinata(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		pinataRequester,
		ipfsCheckRequester,
		agentCollectionRequester,
		agentCollectionsManagerRequester,
	)
	return &WhenPushingImageOfAgentImageProposalTestSuite{
		sut:                              sut,
		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
	}
}

func TestWhenPushingImageOfAgentImageProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image proposal image", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)
			suite.agentCollectionRequester.EXPECT().GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

			err := suite.sut.PushImageOfAgentImageProposal(testChainId, common.HexToAddress(""), *big.NewInt(123))

			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)
			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testCid, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(assert.AnError).Times(2)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			err := suite.sut.PushImageOfAgentImageProposal(testChainId, common.HexToAddress(""), *big.NewInt(123))

			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)

			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testCid, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			err := suite.sut.PushImageOfAgentImageProposal(testChainId, common.HexToAddress(""), *big.NewInt(123))

			assert.Equal(t, err, nil)
		})
	})
}
