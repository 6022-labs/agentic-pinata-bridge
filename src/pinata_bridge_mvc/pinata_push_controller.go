package pinata_bridge_mvc

import (
	"math/big"
	"strings"

	"github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
)

const (
	PUSH_MISSING_IMAGE_CIDS      = "/push_missing_image_cids"
	PUSH_MISSING_IMAGES_OF_AGENT = "/push_missing_images_of_agent/:agentCollectionAddress/:agentCollectionTokenId"
	PUSH_IMAGES_OF_MINT_PROPOSAL = "/push_images_of_mint_proposal/:agentCollectionAddress/:mintProposalId"
)

type PinataPushController struct {
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface
}

func NewPinataPushController(
	pushAgentImageCidToPinata use_cases.PushAgentImageCidToPinataInterface,
) *PinataPushController {
	return &PinataPushController{
		pushAgentImageCidToPinata: pushAgentImageCidToPinata,
	}
}

func (controller *PinataPushController) InitRoutes(c fiber.Router) {
	c.Post(PUSH_MISSING_IMAGE_CIDS, controller.PushMissingImageCids)
	c.Post(PUSH_MISSING_IMAGES_OF_AGENT, controller.PushMissingImagesOfAgent)
	c.Post(PUSH_IMAGES_OF_MINT_PROPOSAL, controller.PushImagesOfMintProposal)
}

func (controller *PinataPushController) PushMissingImageCids(c *fiber.Ctx) error {
	return controller.pushAgentImageCidToPinata.PushMissingImageCids()
}

func (controller *PinataPushController) PushMissingImagesOfAgent(c *fiber.Ctx) error {
	agentCollectionAddressStr := c.Params("agentCollectionAddress")
	if len(strings.TrimSpace(agentCollectionAddressStr)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionAddress is required")
	}

	agentCollectionTokenIdStr := c.Params("agentCollectionTokenId")
	if len(strings.TrimSpace(agentCollectionTokenIdStr)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionTokenId is required")
	}

	if !common.IsHexAddress(agentCollectionAddressStr) {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionAddress is invalid")
	}

	agentCollectionAddress := common.HexToAddress(agentCollectionAddressStr)

	tokenId, ok := big.NewInt(0).SetString(agentCollectionTokenIdStr, 10)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionTokenId is invalid")
	}

	return controller.pushAgentImageCidToPinata.PushMissingImagesOfAgent(agentCollectionAddress, *tokenId)
}

func (controller *PinataPushController) PushImagesOfMintProposal(c *fiber.Ctx) error {
	agentCollectionAddressStr := c.Params("agentCollectionAddress")
	if len(strings.TrimSpace(agentCollectionAddressStr)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionAddress is required")
	}

	mintProposalIdStr := c.Params("mintProposalId")
	if len(strings.TrimSpace(mintProposalIdStr)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "mintProposalId is required")
	}

	if !common.IsHexAddress(agentCollectionAddressStr) {
		return fiber.NewError(fiber.StatusBadRequest, "agentCollectionAddress is invalid")
	}

	agentCollectionAddress := common.HexToAddress(agentCollectionAddressStr)

	mintProposalId, ok := big.NewInt(0).SetString(mintProposalIdStr, 10)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "mintProposalId is invalid")
	}

	return controller.pushAgentImageCidToPinata.PushImagesOfMintProposal(
		agentCollectionAddress,
		*mintProposalId,
	)
}
