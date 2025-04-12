package pinata_bridge_mvc_api

import (
	"math/big"
	"strings"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/gofiber/fiber/v2"
)

const (
	PUSH_ALL_AGENT_IMAGE_CIDS = "/push_all_agent_image_cids"
	PUSH_FROM_AGENT_TOKEN_ID  = "/push_from_agent_token_id/:tokenId"
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
	c.Post(PUSH_FROM_AGENT_TOKEN_ID, controller.PushFromAgentTokenId)
	c.Post(PUSH_ALL_AGENT_IMAGE_CIDS, controller.PushAllAgentImageCids)
}

func (controller *PinataPushController) PushAllAgentImageCids(c *fiber.Ctx) error {
	return controller.pushAgentImageCidToPinata.PushAllAgentImageCids()
}

func (controller *PinataPushController) PushFromAgentTokenId(c *fiber.Ctx) error {
	tokenIdStr := c.Params("tokenId")
	if len(strings.TrimSpace(tokenIdStr)) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "tokenId is required")
	}

	tokenId, ok := big.NewInt(0).SetString(tokenIdStr, 10)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "tokenId is invalid")
	}

	return controller.pushAgentImageCidToPinata.PushFromAgentTokenId(*tokenId)
}
