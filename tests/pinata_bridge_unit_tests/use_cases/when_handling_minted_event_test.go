package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/tests/common_tests/model_builders"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenHandlingMintedEventTestSuite struct {
	sut *use_cases.HandleMintedEvent

	agentCollectionRequester *interfaces_mocks.MockAgentCollectionRequesterInterface
	pinMetrics               *metrics_mocks.MockPinMetricsInterface
}

func WhenHandlingMintedEventBeforeEach(t *testing.T) *WhenHandlingMintedEventTestSuite {
	mockController := gomock.NewController(t)

	pinataRequester := interfaces_mocks.NewMockPinataRequesterInterface(mockController)
	ipfsCheckRequester := interfaces_mocks.NewMockIpfsCheckRequesterInterface(mockController)
	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	// The push use case runs for real so the event's chain id and token id are asserted where they land.
	pinTracer := newNoopPinTracer(mockController)

	pushMissingImagesOfAgent := use_cases.NewPushMissingImagesOfAgent(
		zap.NewNop(),
		services.NewCidPinner(zap.NewNop(), pinataRequester, ipfsCheckRequester, pinMetrics, pinTracer),
		agentCollectionRequester,
		pinataRequester,
		pinMetrics,
	)

	return &WhenHandlingMintedEventTestSuite{
		sut: use_cases.NewHandleMintedEvent(pushMissingImagesOfAgent),

		agentCollectionRequester: agentCollectionRequester,
		pinMetrics:               pinMetrics,
	}
}

func TestWhenHandlingMintedEvent(t *testing.T) {
	t.Parallel()

	t.Run("Given a minted event", func(t *testing.T) {
		t.Parallel()

		t.Run("Should push to pinata using agent token id", func(t *testing.T) {
			t.Parallel()

			suite := WhenHandlingMintedEventBeforeEach(t)

			suite.agentCollectionRequester.EXPECT().
				GetAgentImages(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(
					_ context.Context,
					chainId uint64,
					collectionAddress common.Address,
					collectionAgentTokenId big.Int,
				) ([]string, error) {
					assert.Equal(t, uint64(80002), chainId)
					assert.Equal(t, "123", collectionAgentTokenId.String())
					assert.Equal(t, "0x1234567890123456789012345678901234567890", collectionAddress.Hex())
					return nil, nil
				})
			suite.pinMetrics.EXPECT().
				RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAgent, gomock.Any(), false)

			err := suite.sut.Execute(context.Background(), 80002, model_builders.NewMintedEventBuilder().Build())

			assert.NoError(t, err)
		})
	})
}
