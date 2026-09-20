package services_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/common_tests/model_builders"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenHandlingMintedEventTestingSuite struct {
	sut *services.MintedEventHandler

	agentImagesPinner *interfaces_mocks.MockAgentImagesPinnerInterface
}

func WhenHandlingMintedEventBeforeEach(t *testing.T) *WhenHandlingMintedEventTestingSuite {
	mockController := gomock.NewController(t)
	agentImagesPinner := interfaces_mocks.NewMockAgentImagesPinnerInterface(mockController)

	return &WhenHandlingMintedEventTestingSuite{
		sut:               services.NewMintedEventHandler(agentImagesPinner),
		agentImagesPinner: agentImagesPinner,
	}
}

func TestWhenHandlingMintedEvent(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x2222222222222222222222222222222222222222")

	t.Run("Given an event from a watched collection", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenHandlingMintedEventTestingSuite) {
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), testChainId, collectionAddress, *big.NewInt(456)).
				Return(nil)
		}

		t.Run("Should pin using the event's collection and id", func(t *testing.T) {
			t.Parallel()

			suite := WhenHandlingMintedEventBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Handle(context.Background(), testChainId, model_builders.NewMintedEventBuilder().
				WithCollectionAddress(collectionAddress).
				WithTokenId(big.NewInt(456)).
				Build())

			assert.NoError(t, err)
		})
	})

	t.Run("Given the agentImagesPinner fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenHandlingMintedEventTestingSuite) {
			suite.agentImagesPinner.EXPECT().
				PinMissing(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError)
		}

		t.Run("Should surface the error to the listener", func(t *testing.T) {
			t.Parallel()

			suite := WhenHandlingMintedEventBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Handle(context.Background(), testChainId, model_builders.NewMintedEventBuilder().Build())

			assert.Equal(t, assert.AnError, err)
		})
	})
}
