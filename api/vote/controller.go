package vote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiluk/pubg-heat-drop/models"
)

type VoteController struct {
	service *VoteService
}

func NewController(service *VoteService) *VoteController {
	return &VoteController{
		service: service,
	}
}

func (c *VoteController) PostVote(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if len(authHeader) < 8 {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid or missing Authorization header")
	}
	sessionID := authHeader[7:]
	_, err := uuid.Parse(sessionID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid session ID")
	}

	var heat models.Heat
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
