package pinata_bridge_mvc

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/common/mvc"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases/requests"
	"github.com/gofiber/fiber/v3"
)

const (
	PushMissingImageCidsEndpoint     = "/push_missing_image_cids"
	PushMissingImagesOfAgentEndpoint = "/push_missing_images_of_agent" +
		"/:chainId/:agentCollectionAddress/:agentCollectionTokenId"
	PushImagesOfMintProposalEndpoint = "/push_images_of_mint_proposal/:chainId/:agentCollectionAddress/:mintProposalId"
)

type PinataPushController struct {
	pushMissingImageCids     *use_cases.PushMissingImageCids
	pushMissingImagesOfAgent *use_cases.PushMissingImagesOfAgent
	pushImagesOfMintProposal *use_cases.PushImagesOfMintProposal
	errorWriter              *mvc.ErrorWriter
}

func NewPinataPushController(
	pushMissingImageCids *use_cases.PushMissingImageCids,
	pushMissingImagesOfAgent *use_cases.PushMissingImagesOfAgent,
	pushImagesOfMintProposal *use_cases.PushImagesOfMintProposal,
	errorWriter *mvc.ErrorWriter,
) *PinataPushController {
	return &PinataPushController{
		pushMissingImageCids:     pushMissingImageCids,
		pushMissingImagesOfAgent: pushMissingImagesOfAgent,
		pushImagesOfMintProposal: pushImagesOfMintProposal,
		errorWriter:              errorWriter,
	}
}

func (c *PinataPushController) RegisterRoutes(app fiber.Router) {
	app.Post(PushMissingImageCidsEndpoint, c.PushMissingImageCids)
	app.Post(PushMissingImagesOfAgentEndpoint, c.PushMissingImagesOfAgent)
	app.Post(PushImagesOfMintProposalEndpoint, c.PushImagesOfMintProposal)
}

func (c *PinataPushController) PushMissingImageCids(ctx fiber.Ctx) error {
	if _, err := c.pushMissingImageCids.Execute(ctx.Context()); err != nil {
		return c.errorWriter.Write(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *PinataPushController) PushMissingImagesOfAgent(ctx fiber.Ctx) error {
	request := &requests.PushMissingImagesOfAgentRequest{}
	if err := ctx.Bind().URI(request); err != nil {
		return c.errorWriter.Write(ctx, errors.NewBadRequestError("invalid_params", err.Error()))
	}

	if _, err := c.pushMissingImagesOfAgent.Execute(ctx.Context(), request); err != nil {
		return c.errorWriter.Write(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *PinataPushController) PushImagesOfMintProposal(ctx fiber.Ctx) error {
	request := &requests.PushImagesOfMintProposalRequest{}
	if err := ctx.Bind().URI(request); err != nil {
		return c.errorWriter.Write(ctx, errors.NewBadRequestError("invalid_params", err.Error()))
	}

	if _, err := c.pushImagesOfMintProposal.Execute(ctx.Context(), request); err != nil {
		return c.errorWriter.Write(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
