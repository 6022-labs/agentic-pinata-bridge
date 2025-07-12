package pinata_bridge_mvc_api

import (
	"math/big"
	"strings"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
)

const (
	PUSH_ALL_AGENT_IMAGE_CIDS    = "/push_all_agent_image_cids"
	PUSH_IMAGE_OF_AGENT          = "/push_image_of_agent/:agentCollectionAddress/:agentCollectionTokenId"
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
	c.Post(PUSH_IMAGE_OF_AGENT, controller.PushImagesOfAgent)
	c.Post(PUSH_ALL_AGENT_IMAGE_CIDS, controller.PushAllAgentImageCids)
	c.Post(PUSH_IMAGES_OF_MINT_PROPOSAL, controller.PushImagesOfMintProposal)
}

func (controller *PinataPushController) PushAllAgentImageCids(c *fiber.Ctx) error {
	return controller.pushAgentImageCidToPinata.PushAllAgentImageCids()
}

func (controller *PinataPushController) PushImagesOfAgent(c *fiber.Ctx) error {
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

	return controller.pushAgentImageCidToPinata.PushImagesOfAgent(agentCollectionAddress, *tokenId)
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
