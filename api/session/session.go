package session

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Session struct {
	SessionID string `json:"sessionID"`
	HasVoted  bool   `json:"hasVoted"`
	Lobby     string `json:"lobby"`
}

type SessionHandler struct {
	cache *redis.Client
}

func NewSessionHandler(cache *redis.Client) *SessionHandler {
	return &SessionHandler{
		cache: cache,
	}
}

func (sh *SessionHandler) CreateSession(ctx *fiber.Ctx) error {
	session := Session{
		SessionID: uuid.NewString(),
		HasVoted:  false,
		Lobby:     "",
	}

	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = sh.cache.Set(ctx.Context(), "session:"+session.SessionID, sessionJSON, 0).Err()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	log.Println("Session created with ID:", session.SessionID)

	return ctx.JSON(session)
}

func (sh *SessionHandler) GetSession(ctx *fiber.Ctx) error {
	sessionID := ctx.Params("id")
	if sessionID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing session ID")
	}

	log.Println("Getting session with ID:", sessionID)

	sessionJSON, err := sh.cache.Get(ctx.Context(), "session:"+sessionID).Result()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Session not found")
	}

	var session Session
	err = json.Unmarshal([]byte(sessionJSON), &session)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(session)
}
