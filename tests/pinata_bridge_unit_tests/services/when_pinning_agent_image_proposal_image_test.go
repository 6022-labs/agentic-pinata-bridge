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

type WhenPinningAgentImageProposalImageTestSuite struct {
	sut *services.AgentImageProposalImagePinner

	pinataRequester          *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester       *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester *interfaces_mocks.MockAgentCollectionRequesterInterface
	pinMetrics               *metrics_mocks.MockPinMetricsInterface
}

func WhenPinningAgentImageProposalImageBeforeEach(t *testing.T) *WhenPinningAgentImageProposalImageTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	pinTracer := newNoopPinTracer(mockController)

	cidPinner := services.NewCidPinner(zap.NewNop(), pinataRequester, ipfsCheckRequester, pinMetrics, pinTracer)

	sut := services.NewAgentImageProposalImagePinner(
		zap.NewNop(),
		cidPinner,
		agentCollectionRequester,
		pinMetrics,
	)
	return &WhenPinningAgentImageProposalImageTestSuite{
		sut:                      sut,
		pinataRequester:          pinataRequester,
		ipfsCheckRequester:       ipfsCheckRequester,
		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func TestWhenPinningAgentImageProposalImage(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting agent image proposal image", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningAgentImageProposalImageTestSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindImageProposal, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningAgentImageProposalImageBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "image_proposal_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid

		initSuite := func(suite *WhenPinningAgentImageProposalImageTestSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(&testCid, nil)
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
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomeFailed)
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindImageProposal, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningAgentImageProposalImageBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(
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

	t.Run("Given the proposal image is not a cid", func(t *testing.T) {
		t.Parallel()

		nonCidImage := testNonCidImage

		initSuite := func(suite *WhenPinningAgentImageProposalImageTestSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(&nonCidImage, nil)

			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomeInvalidCid)
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindImageProposal, gomock.Any(), false)
		}

		t.Run("Should skip it without calling pinata", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningAgentImageProposalImageBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			assert.Nil(t, err)
		})
	})

	t.Run("Given no error occurs", func(t *testing.T) {
		t.Parallel()

		testCid := testValidCid

		initSuite := func(suite *WhenPinningAgentImageProposalImageTestSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetAgentImageProposalImage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(&testCid, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().
				GetHostNodeIds(gomock.Any(), gomock.Any()).
				Return([]string{testHostNodeId}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomePinned)
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindImageProposal, gomock.Any(), false)
		}

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningAgentImageProposalImageBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Pin(
				context.Background(),
				testChainId,
				common.HexToAddress(testCollectionAddress),
				*big.NewInt(123),
			)

			assert.Equal(t, err, nil)
		})
	})
}
