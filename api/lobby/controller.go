package lobby

import (
	"github.com/gofiber/fiber/v2"
)

type LobbyController struct {
	service *LobbyService
}

func NewController(service *LobbyService) *LobbyController {
	return &LobbyController{
		service: service,
	}
}

func (c *LobbyController) PostLobby(ctx *fiber.Ctx) error {
	lobby, err := c.service.CreateLobby(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(lobby)
}

func (c *LobbyController) GetLobby(ctx *fiber.Ctx) error {
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
