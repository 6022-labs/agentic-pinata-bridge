package use_cases_test

import (
	"context"
	"math/big"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPushingImageOfAgentImageProposalTestingSuite struct {
	sut *use_cases.PushImageOfAgentImageProposal

	agentImageProposalImagePinner *interfaces_mocks.MockAgentImageProposalImagePinnerInterface
}

func WhenPushingImageOfAgentImageProposalBeforeEach(t *testing.T) *WhenPushingImageOfAgentImageProposalTestingSuite {
	mockController := gomock.NewController(t)
	agentImageProposalImagePinner := interfaces_mocks.NewMockAgentImageProposalImagePinnerInterface(mockController)

	return &WhenPushingImageOfAgentImageProposalTestingSuite{
		sut:                           use_cases.NewPushImageOfAgentImageProposal(agentImageProposalImagePinner),
		agentImageProposalImagePinner: agentImageProposalImagePinner,
	}
}

func TestWhenPushingImageOfAgentImageProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid request", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingImageOfAgentImageProposalTestingSuite) {
			suite.agentImageProposalImagePinner.EXPECT().
				Pin(gomock.Any(), testChainId, common.HexToAddress(testCollectionAddress), *big.NewInt(123)).
				Return(nil)
		}

		t.Run("Should hand the parsed values to the agentImageProposalImagePinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)
			initSuite(suite)

			response, err := suite.sut.Execute(context.Background(), &requests.PushImageOfAgentImageProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				AgentImageProposalId: "123",
			})

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	})

	t.Run("Given the agentImageProposalImagePinner fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingImageOfAgentImageProposalTestingSuite) {
			suite.agentImageProposalImagePinner.EXPECT().
				Pin(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(apperrors.NewUnavailableError("image_proposal_read_failed", "upstream request failed"))
		}

		t.Run("Should return its error untouched", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)
			initSuite(suite)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImageOfAgentImageProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				AgentImageProposalId: "123",
			})

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "image_proposal_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given an empty request", func(t *testing.T) {
		t.Parallel()

		t.Run("Should reject the request before touching the agentImageProposalImagePinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImageOfAgentImageProposalBeforeEach(t)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImageOfAgentImageProposalRequest{})

			var validationError *apperrors.ValidationError
			assert.ErrorAs(t, err, &validationError)
		})
	})
}
