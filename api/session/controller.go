package session

import "github.com/gofiber/fiber/v2"

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) PostSession(ctx *fiber.Ctx) error {
	session, err := c.service.CreateSession(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(session)
}

func (c *Controller) GetSession(ctx *fiber.Ctx) error {
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
