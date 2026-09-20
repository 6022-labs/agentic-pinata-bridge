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

type WhenHandlingAgentImageProposalCreatedEventTestingSuite struct {
	sut *services.AgentImageProposalCreatedEventHandler

	agentImageProposalImagePinner *interfaces_mocks.MockAgentImageProposalImagePinnerInterface
}

func WhenHandlingAgentImageProposalCreatedEventBeforeEach(
	t *testing.T,
) *WhenHandlingAgentImageProposalCreatedEventTestingSuite {
	mockController := gomock.NewController(t)
	agentImageProposalImagePinner := interfaces_mocks.NewMockAgentImageProposalImagePinnerInterface(mockController)

	return &WhenHandlingAgentImageProposalCreatedEventTestingSuite{
		sut:                           services.NewAgentImageProposalCreatedEventHandler(agentImageProposalImagePinner),
		agentImageProposalImagePinner: agentImageProposalImagePinner,
	}
}

func TestWhenHandlingAgentImageProposalCreatedEvent(t *testing.T) {
	t.Parallel()

	collectionAddress := common.HexToAddress("0x2222222222222222222222222222222222222222")

	t.Run("Given an event from a watched collection", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenHandlingAgentImageProposalCreatedEventTestingSuite) {
			suite.agentImageProposalImagePinner.EXPECT().
				Pin(gomock.Any(), testChainId, collectionAddress, *big.NewInt(456)).
				Return(nil)
		}

		t.Run("Should pin using the event's collection and id", func(t *testing.T) {
			t.Parallel()

			suite := WhenHandlingAgentImageProposalCreatedEventBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Handle(
				context.Background(),
				testChainId,
				model_builders.NewAgentImageProposalCreatedEventBuilder().
					WithCollectionAddress(collectionAddress).
					WithProposalId(big.NewInt(456)).
					Build(),
			)

			assert.NoError(t, err)
		})
	})

	t.Run("Given the agentImageProposalImagePinner fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenHandlingAgentImageProposalCreatedEventTestingSuite) {
			suite.agentImageProposalImagePinner.EXPECT().
				Pin(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(assert.AnError)
		}

		t.Run("Should surface the error to the listener", func(t *testing.T) {
			t.Parallel()

			suite := WhenHandlingAgentImageProposalCreatedEventBeforeEach(t)
			initSuite(suite)

			err := suite.sut.Handle(
				context.Background(),
				testChainId,
				model_builders.NewAgentImageProposalCreatedEventBuilder().Build(),
			)

			assert.Equal(t, assert.AnError, err)
		})
	})
}
