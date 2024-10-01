package session

import "github.com/gofiber/fiber/v2"

type SessionController struct {
	service *SessionService
}

func NewController(service *SessionService) *SessionController {
	return &SessionController{
		service: service,
	}
}

func (c *SessionController) PostSession(ctx *fiber.Ctx) error {
	session, err := c.service.CreateSession(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(session)
}

func (c *SessionController) GetSession(ctx *fiber.Ctx) error {
	sessionID := ctx.Params("id")
	if sessionID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing session ID")
	}

	session, err := c.service.GetSession(ctx, sessionID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(session)
}
