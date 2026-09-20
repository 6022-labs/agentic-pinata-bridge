package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	metrics_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/metrics_mocks/interfaces_mocks"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type WhenPushingMissingImageCidsTestingSuite struct {
	sut *use_cases.PushMissingImageCids

	agentCollectionRequester         *interfaces_mocks.MockAgentCollectionRequesterInterface
	agentCollectionsManagerRequester *interfaces_mocks.MockAgentCollectionsManagerRequesterInterface
	agentImagesPinner                *interfaces_mocks.MockAgentImagesPinnerInterface
	pinMetrics                       *metrics_mocks.MockPinMetricsInterface
}

func WhenPushingMissingImageCidsBeforeEach(t *testing.T) *WhenPushingMissingImageCidsTestingSuite {
	mockController := gomock.NewController(t)

	agentCollectionRequester := interfaces_mocks.NewMockAgentCollectionRequesterInterface(mockController)
	agentCollectionsManagerRequester := interfaces_mocks.NewMockAgentCollectionsManagerRequesterInterface(
		mockController,
	)
	agentImagesPinner := interfaces_mocks.NewMockAgentImagesPinnerInterface(mockController)
	pinMetrics := metrics_mocks.NewMockPinMetricsInterface(mockController)

	sut := use_cases.NewPushMissingImageCids(
		zap.NewNop(),
		agentCollectionRequester,
		pinMetrics,
		settings.NewChainsSettingsFromChainIds([]uint64{testChainId}),
		agentCollectionsManagerRequester,
		agentImagesPinner,
		newNoopPinTracer(mockController),
	)

	return &WhenPushingMissingImageCidsTestingSuite{
		sut: sut,

		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		agentImagesPinner:                agentImagesPinner,
		pinMetrics:                       pinMetrics,
	}
}

func TestWhenPushingMissingImageCids(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	t.Run("Given error occurs while getting all collections addresses", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImageCidsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return(nil, assert.AnError)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background())

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "collections_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given error occurs while getting all token ids", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingMissingImageCidsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{collectionAddress}, nil)
			suite.agentCollectionRequester.EXPECT().
				GetAllTokenIds(gomock.Any(), testChainId, collectionAddress).
				Return(nil, assert.AnError)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), true)
		}

		t.Run("Should return error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background())

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "token_ids_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given every collection and token can be listed", func(t *testing.T) {
		t.Parallel()

		tokenIds := []big.Int{*big.NewInt(1), *big.NewInt(2)}

		initSuite := func(suite *WhenPushingMissingImageCidsTestingSuite) {
			suite.agentCollectionsManagerRequester.EXPECT().
				GetAllCollectionAddresses(gomock.Any(), testChainId).
				Return([]common.Address{collectionAddress}, nil)
			suite.agentCollectionRequester.EXPECT().
				GetAllTokenIds(gomock.Any(), testChainId, collectionAddress).
				Return(tokenIds, nil)
			// One agent failing must not stop the sweep.
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), testChainId, collectionAddress, tokenIds[0]).
				Return(assert.AnError)
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), testChainId, collectionAddress, tokenIds[1]).
				Return(nil)
			suite.pinMetrics.EXPECT().RecordSweep(gomock.Any(), metrics_interfaces.SweepKindAll, gomock.Any(), false)
		}

		t.Run("Should sweep every agent and report success", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingMissingImageCidsBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background())

			assert.NoError(t, err)
		})
	})
}
