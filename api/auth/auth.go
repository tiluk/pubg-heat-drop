package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/tiluk/pubg-heat-drop/models"
	"github.com/tiluk/pubg-heat-drop/session"
)

func NewAuthMiddleware(sessionService *session.SessionService) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		if ctx.Path() == "/api/auth" {
			return ctx.Next()
		}

		authHeader := ctx.Get("Authorization")
		if len(authHeader) < 8 && authHeader[0:7] != "Bearer " {
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid or missing Authorization header")
		}

		rawJWT := authHeader[7:]
		if rawJWT == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		token, err := jwt.Parse(rawJWT, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		if !token.Valid {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		sessionID, ok := claims["sessionID"].(string)
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		isValidSession, err := sessionService.VerifyJWTSession(ctx, &models.Session{
			SessionID: sessionID,
		})
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		if !isValidSession {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("sessionID", sessionID)

		hasVoted, err := sessionService.GetHasVoted(ctx, sessionID)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		ctx.Locals("hasVoted", hasVoted)

		return ctx.Next()
	}
}
