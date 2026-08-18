package services_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/services_mocks/interfaces_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenHandlingMintedEventTestSuite struct {
	sut *services.MintedEventHandler

	pushAgentImageCidToPinata *interfaces_mocks.MockPushAgentImageCidToPinataInterface
}

func WhenHandlingMintedEventBeforeEach(t *testing.T) *WhenHandlingMintedEventTestSuite {
	mockController := gomock.NewController(t)

	pushAgentImageCidToPinata := interfaces_mocks.NewMockPushAgentImageCidToPinataInterface(mockController)

	sut := services.NewMintedEventHandler(pushAgentImageCidToPinata)

	return &WhenHandlingMintedEventTestSuite{
		sut: sut,

		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func TestWhenHandlingMintedEvent(t *testing.T) {
	t.Parallel()

	t.Run("Should push to pinata using agent token id", func(t *testing.T) {
		t.Parallel()

		suite := WhenHandlingMintedEventBeforeEach(t)
		suite.pushAgentImageCidToPinata.EXPECT().PushMissingImagesOfAgent(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, chainId uint64, collectionAddress common.Address, collectionAgentTokenId big.Int) error {
				assert.Equal(t, chainId, uint64(80002))
				assert.Equal(t, collectionAgentTokenId.String(), "123")
				return nil
			})

		err := suite.sut.Handle(context.Background(), 80002, &abi.AgentCollectionV1Minted{
			Raw: types.Log{
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
			TokenId: big.NewInt(123),
		})

		assert.Equal(t, err, nil)
	})
}
