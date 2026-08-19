package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

const testChainId uint64 = 80002

const (
	testChainIdString     = "80002"
	testCollectionAddress = "0x0000000000000000000000000000000000000000"
)

type WhenPushingMissingImageCidsTestingSuite struct {
	sut *use_cases.PushMissingImageCids

	pinataRequester                  *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	pinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingMissingImageCidsBeforeEach(t *testing.T) *WhenPushingMissingImageCidsTestingSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
		mockController,
	)

	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	pinTracer := newNoopPinTracer(mockController)

	cidPinner := services.NewCidPinner(zap.NewNop(), pinataRequester, ipfsCheckRequester, pinMetrics, pinTracer)

	pushMissingImagesOfAgent := use_cases.NewPushMissingImagesOfAgent(
		zap.NewNop(),
		cidPinner,
		agentCollectionRequester,
		pinataRequester,
		pinMetrics,
	)

	sut := use_cases.NewPushMissingImageCids(
		zap.NewNop(),
		agentCollectionRequester,
		pinMetrics,
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		agentCollectionsManagerRequester,
		pushMissingImagesOfAgent,
		pinTracer,
	)
	return &WhenPushingMissingImageCidsTestingSuite{
		sut: sut,

		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pinMetrics:                       pinMetrics,
	}
}

func TestWhenPushingMissingImageCids(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting all collections addresses", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImageCidsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background())
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
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), gomock.Any()).
				Return(collectionAddress, nil)
			for _, address := range collectionAddress {
				suite.agentCollectionRequester.EXPECT().
					GetAllTokenIds(gomock.Any(), gomock.Any(), address).
					Return(nil, assert.AnError)
			}

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), true)

			_, err := suite.sut.Execute(context.Background())
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

			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), gomock.Any()).
				Return(collectionAddress, nil)
			for _, address := range collectionAddress {
				suite.agentCollectionRequester.EXPECT().
					GetAllTokenIds(gomock.Any(), gomock.Any(), address).
					Return(tokenIds, nil)
				for _, tokenId := range tokenIds {
					suite.agentCollectionRequester.EXPECT().
						GetAgentImages(gomock.Any(), gomock.Any(), address, tokenId).
						Return([]string{imageCid}, nil)
				}
			}

			isCidUploaded := false

			suite.pinataRequester.EXPECT().
				IsCidUploaded(gomock.Any(), imageCid).
				Return(&isCidUploaded, nil).
				Times(len(tokenIds))
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil).
				Times(len(tokenIds))
			suite.ipfsCheckRequester.EXPECT().GetMultiAddresses(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeEmpty, int64(3)).
				Times(len(tokenIds))
			suite.pinMetrics.EXPECT().
				RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, false, gomock.Any()).
				Times(len(tokenIds))
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned).
				Times(len(tokenIds))
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false).
				Times(len(tokenIds))
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), false)

			_, err := suite.sut.Execute(context.Background())
			assert.NoError(t, err)
		})
	})
}
