package vote

import "github.com/gofiber/fiber/v2"

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) PostVote(ctx *fiber.Ctx) error {
	lobbyID := ctx.Params("lobby-id")
	if lobbyID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing lobby ID")
	}

	heatmap, err := c.service.CastVote(lobbyID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(heatmap)

}
