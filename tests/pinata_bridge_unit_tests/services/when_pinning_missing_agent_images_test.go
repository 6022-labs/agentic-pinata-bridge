package services_test

import (
	"context"
	"math/big"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPinningMissingAgentImagesTestingSuite struct {
	sut *services.AgentImagesPinner

	pinataRequester          *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester       *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester *interfaces_mocks.MockAgentCollectionRequesterInterface
	pinMetrics               *metrics_mocks.MockPinMetricsInterface
}

func WhenPinningMissingAgentImagesBeforeEach(t *testing.T) *WhenPinningMissingAgentImagesTestingSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	pinTracer := newNoopPinTracer(mockController)

	cidPinner := services.NewCidPinner(zap.NewNop(), pinataRequester, ipfsCheckRequester, pinMetrics, pinTracer)

	sut := services.NewAgentImagesPinner(
		zap.NewNop(),
		cidPinner,
		agentCollectionRequester,
		pinataRequester,
		pinMetrics,
	)

	return &WhenPinningMissingAgentImagesTestingSuite{
		sut: sut,

		pinataRequester:          pinataRequester,
		ipfsCheckRequester:       ipfsCheckRequester,
		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func TestWhenPinningMissingAgentImages(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "agent_images_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given error occurs while checking if cid is already pinned", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			tokenId := big.NewInt(123)
			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*tokenId,
			)

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "pin_status_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid
		isCidUploaded := false

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError).
				Times(2)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return([]string{testHostNodeId}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			tokenId := big.NewInt(123)
			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*tokenId,
			)

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "image_pin_failed", unavailableError.Code)
		})
	})

	t.Run("Given an agent image is not a cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testNonCidImage}, nil)

			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeInvalidCid)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)
		}

		t.Run("Should skip it without calling pinata", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			assert.Nil(t, err)
		})
	})

	t.Run("Given an agent image fails to pin", func(t *testing.T) {
		t.Parallel()

		isCidUploaded := false

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testValidCid, testOtherValidCid}, nil)
			suite.pinataRequester.EXPECT().
				IsCidUploaded(gomock.Any(), gomock.Any()).
				Return(&isCidUploaded, nil).
				Times(2)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError).
				AnyTimes()

			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), testValidCid, nil).Return(assert.AnError)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), testOtherValidCid, nil).Return(nil)

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFailed, gomock.Any()).
				Times(2)
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, false, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), true)
		}

		t.Run("Should still pin the remaining images and report the failure", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "image_pin_failed", unavailableError.Code)
		})
	})

	t.Run("Given image cid is already pinned", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid
		isCidUploaded := true

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)

			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeAlreadyPinned)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)
		}

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			tokenId := big.NewInt(123)
			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*tokenId,
			)

			assert.Nil(t, err)
		})
	})

	t.Run("Given image cid is not yet pinned", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid
		isCidUploaded := false

		initSuite := func(suite *WhenPinningMissingAgentImagesTestingSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().IsCidUploaded(gomock.Any(), testCid).Return(&isCidUploaded, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return([]string{testHostNodeId}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)
		}

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningMissingAgentImagesBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinMissing(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			assert.Nil(t, err)
		})
	})
}
