package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingImagesOfMintProposalTestSuite struct {
	sut *use_cases.PushImagesOfMintProposal

	pinataRequester                  *interfaces_mocks.MockPinataRequesterInterface
	ipfsCheckRequester               *interfaces_mocks.MockIpfsCheckRequesterInterface
	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	pinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingImagesOfMintProposalBeforeEach(t *testing.T) *WhenPushingImagesOfMintProposalTestSuite {
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

	sut := use_cases.NewPushImagesOfMintProposal(
		zap.NewNop(),
		cidPinner,
		agentCollectionRequester,
		pinMetrics,
	)
	return &WhenPushingImagesOfMintProposalTestSuite{
		sut:                              sut,
		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pinMetrics:                       pinMetrics,
	}
}

func TestWhenPushingImagesOfMintProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given error occurs while getting mint proposal images", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingImagesOfMintProposalTestSuite) {
			suite.agentCollectionRequester.EXPECT().
				GetMintProposalImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, assert.AnError)

			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindMintProposal, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImagesOfMintProposalBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImagesOfMintProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				MintProposalId: big.NewInt(123).String(),
			})

			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given error occurs while pushing to pinata", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImagesOfMintProposalBeforeEach(t)

			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().
				GetMintProposalImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().
				PinCid(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError).
				Times(2)
			suite.ipfsCheckRequester.EXPECT().
				GetMultiAddresses(gomock.Any(), gomock.Any()).
				Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, true, gomock.Any())
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomeFailed, false, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindMintProposal, metrics_interfaces.PinOutcomeFailed)
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindMintProposal, gomock.Any(), true)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImagesOfMintProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				MintProposalId: big.NewInt(123).String(),
			})

			assert.Equal(t, err, assert.AnError)
		})
	})

	t.Run("Given no error occurs", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImagesOfMintProposalBeforeEach(t)

			testCid := "test-cid"
			suite.agentCollectionRequester.EXPECT().
				GetMintProposalImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return([]string{testCid}, nil)
			suite.pinataRequester.EXPECT().PinCid(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			suite.ipfsCheckRequester.EXPECT().
				GetMultiAddresses(gomock.Any(), gomock.Any()).
				Return([]string{"/ip4/127.0.0.1/tcp/4001"}, nil).
				AnyTimes()

			suite.pinMetrics.EXPECT().
				RecordHostLookup(gomock.Any(), metrics_interfaces.HostLookupOutcomeFound, int64(1))
			suite.pinMetrics.EXPECT().RecordPin(gomock.Any(), metrics_interfaces.PinOutcomePinned, true, gomock.Any())
			suite.pinMetrics.EXPECT().
				RecordSweepImage(gomock.Any(), metrics_interfaces.SweepKindMintProposal, metrics_interfaces.PinOutcomePinned)
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindMintProposal, gomock.Any(), false)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImagesOfMintProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				MintProposalId: big.NewInt(123).String(),
			})

			assert.Equal(t, err, nil)
		})
	})
}
