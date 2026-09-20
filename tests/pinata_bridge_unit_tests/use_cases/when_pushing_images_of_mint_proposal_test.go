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

type WhenPushingImagesOfMintProposalTestingSuite struct {
	sut *use_cases.PushImagesOfMintProposal

	mintProposalImagesPinner *interfaces_mocks.MockMintProposalImagesPinnerInterface
}

func WhenPushingImagesOfMintProposalBeforeEach(t *testing.T) *WhenPushingImagesOfMintProposalTestingSuite {
	mockController := gomock.NewController(t)
	mintProposalImagesPinner := interfaces_mocks.NewMockMintProposalImagesPinnerInterface(mockController)

	return &WhenPushingImagesOfMintProposalTestingSuite{
		sut:                      use_cases.NewPushImagesOfMintProposal(mintProposalImagesPinner),
		mintProposalImagesPinner: mintProposalImagesPinner,
	}
}

func TestWhenPushingImagesOfMintProposal(t *testing.T) {
	t.Parallel()

	t.Run("Given a valid request", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingImagesOfMintProposalTestingSuite) {
			suite.mintProposalImagesPinner.EXPECT().
				Pin(gomock.Any(), testChainId, common.HexToAddress(testCollectionAddress), *big.NewInt(123)).
				Return(nil)
		}

		t.Run("Should hand the parsed values to the mintProposalImagesPinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImagesOfMintProposalBeforeEach(t)
			initSuite(suite)

			response, err := suite.sut.Execute(context.Background(), &requests.PushImagesOfMintProposalRequest{
				ProposalRequest: requests.ProposalRequest{
					CollectionRequest: requests.CollectionRequest{
						ChainId:                testChainIdString,
						AgentCollectionAddress: testCollectionAddress,
					},
				},
				MintProposalId: "123",
			})

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	})

	t.Run("Given the mintProposalImagesPinner fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPushingImagesOfMintProposalTestingSuite) {
			suite.mintProposalImagesPinner.EXPECT().
				Pin(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(apperrors.NewUnavailableError("mint_proposal_images_read_failed", "upstream request failed"))
		}

		t.Run("Should return its error untouched", func(t *testing.T) {
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
				MintProposalId: "123",
			})

			var unavailableError *apperrors.UnavailableError
			assert.ErrorAs(t, err, &unavailableError)
			assert.Equal(t, "mint_proposal_images_read_failed", unavailableError.Code)
		})
	})

	t.Run("Given an empty request", func(t *testing.T) {
		t.Parallel()

		t.Run("Should reject the request before touching the mintProposalImagesPinner", func(t *testing.T) {
			t.Parallel()

			suite := WhenPushingImagesOfMintProposalBeforeEach(t)

			_, err := suite.sut.Execute(context.Background(), &requests.PushImagesOfMintProposalRequest{})

			var validationError *apperrors.ValidationError
			assert.ErrorAs(t, err, &validationError)
		})
	})
}
