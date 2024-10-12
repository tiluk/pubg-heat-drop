package session

import (
	"github.com/gofiber/fiber/v2"
)

type SessionController struct {
	service *SessionService
}

func NewController(service *SessionService) *SessionController {
	return &SessionController{
		service: service,
	}
}

func (c *SessionController) PostSession(ctx *fiber.Ctx) error {
	jwt, err := c.service.CreateSession(ctx)
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).SendString(e.Message)
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	ctx.Response().Header.Set("Authorization", *jwt)

	return ctx.SendStatus(fiber.StatusOK)
}
