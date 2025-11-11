package event_handlers_test

import (
	"math/big"
	"testing"

	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/abi"
	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/event_handlers"
	"github.com/6022-labs/agentic-ai-pinata-bridge/tests/pinata_bridge_mocks/use_cases_mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenHandlingMintedEventTestSuite struct {
	sut *event_handlers.MintedEventHandler

	pushAgentImageCidToPinata *use_cases_mocks.MockPushAgentImageCidToPinataInterface
}

func WhenHandlingMintedEventBeforeEach(t *testing.T) *WhenHandlingMintedEventTestSuite {
	mockController := gomock.NewController(t)

	pushAgentImageCidToPinata := use_cases_mocks.NewMockPushAgentImageCidToPinataInterface(mockController)

	sut := event_handlers.NewMintedEventHandler(pushAgentImageCidToPinata)

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
		suite.pushAgentImageCidToPinata.EXPECT().PushMissingImagesOfAgent(gomock.Any(), gomock.Any()).DoAndReturn(func(collectionAddress common.Address, collectionAgentTokenId big.Int) error {
			assert.Equal(t, collectionAgentTokenId.String(), "123")
			return nil
		})

		err := suite.sut.Handle(&abi.AgentCollectionV1Minted{
			Raw: types.Log{
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
			TokenId: big.NewInt(123),
		})

		assert.Equal(t, err, nil)
	})
}
