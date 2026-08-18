package services_test

import (
	"context"
	"math/big"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingMissingImagesOfAgentTestingSuite struct {
	sut *services.PushAgentImageCidToPinata

	pinataRequester                  *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	pinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingMissingImagesOfAgentBeforeEach(t *testing.T) *WhenPushingMissingImagesOfAgentTestingSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(mockController)

	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	sut := services.NewPushAgentImageCidToPinata(
		zap.NewNop(),
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		pinataRequester,
		ipfsCheckRequester,
		agentCollectionRequester,
		agentCollectionsManagerRequester,
		pinMetrics,
	)

	return &WhenPushingMissingImagesOfAgentTestingSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pinMetrics:                       pinMetrics,
	}
}

func TestWhenPushingImagesOfAgent(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)

			err := suite.sut.PushMissingImagesOfAgent(context.Background(), testChainId, common.HexToAddress(""), *big.NewInt(123))

			assert.NotNil(t, err)
		})
	})

	t.Run("Given error occurs while checking if cid is already pinned", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImagesOfAgentBeforeEach(t)

			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(context.Background(), testChainId, common.HexToAddress(""), *tokenId)

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
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError).Times(2)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
			suite.pinMetrics.EXPECT().RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(context.Background(), testChainId, common.HexToAddress(""), *tokenId)

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
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)

			suite.pinMetrics.EXPECT().RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeAlreadyPinned)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)

			tokenId := big.NewInt(123)
			err := suite.sut.PushMissingImagesOfAgent(context.Background(), testChainId, common.HexToAddress(""), *tokenId)

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
			suite.agentCollectionRequester.EXPECT().GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).AnyTimes()

			suite.pinMetrics.EXPECT().RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)

			err := suite.sut.PushMissingImagesOfAgent(context.Background(), testChainId, common.HexToAddress(""), *big.NewInt(123))

			assert.Nil(t, err)
		})
	})
}
