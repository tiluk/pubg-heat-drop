package lobby

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) PostLobby(ctx *fiber.Ctx) error {
	lobby, err := c.service.CreateLobby(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(lobby)
}

func (c *Controller) GetLobby(ctx *fiber.Ctx) error {
	lobbyID := ctx.Params("id")
	if lobbyID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing lobby ID")
	}

	lobby, err := c.service.GetLobby(ctx, lobbyID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(lobby)
}
