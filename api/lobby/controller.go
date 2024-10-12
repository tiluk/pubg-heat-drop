package lobby

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiluk/pubg-heat-drop/models"
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
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).SendString(e.Message)
	}
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
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).SendString(e.Message)
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(lobby)
}

func (c *LobbyController) PostLobbyVote(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if len(authHeader) < 8 && authHeader[0:7] != "Bearer " {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid or missing Authorization header")
	}
	sessionID := ctx.Locals("sessionID").(string)
	_, err := uuid.Parse(sessionID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid session ID")
	}

	lobbyID := ctx.Params("id")
	if lobbyID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing lobby ID")
	}

	var heat models.Heat
	err = ctx.BodyParser(&heat)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.service.AddLobbyVote(ctx, lobbyID, sessionID, heat)
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).SendString(e.Message)
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
