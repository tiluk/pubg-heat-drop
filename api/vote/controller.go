package vote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiluk/pubg-heat-drop/lobby"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) PostVote(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if len(authHeader) < 8 {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid or missing Authorization header")
	}
	sessionID := authHeader[7:]
	_, err := uuid.Parse(sessionID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid session ID")
	}

	var heat lobby.Heat
	err = ctx.BodyParser(&heat)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	lobbyID := ctx.Params("id")
	if lobbyID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing lobby ID")
	}

	heatmap, err := c.service.CastVote(ctx, sessionID, lobbyID, &heat)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(heatmap)
}
