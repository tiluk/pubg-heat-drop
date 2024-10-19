package session

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/tiluk/pubg-heat-drop/models"
)

type SessionRepository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *SessionRepository {
	return &SessionRepository{
		cache: cache,
	}
}

func toSessionKey(sessionID string) string {
	return "session:" + sessionID
}

func (r *SessionRepository) CreateSession(ctx *fiber.Ctx, session *models.Session) error {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toSessionKey(session.SessionID), sessionJSON, 0).Err()
}

func (r *SessionRepository) GetSession(ctx *fiber.Ctx, sessionID string) (*models.Session, error) {
	sessionJSON, err := r.cache.Get(ctx.Context(), toSessionKey(sessionID)).Result()
	if err == redis.Nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "session not found")
	} else if err != nil {
		return nil, err
	}

	var session models.Session
	err = json.Unmarshal([]byte(sessionJSON), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *SessionRepository) SetHasVoted(ctx *fiber.Ctx, sessionID string) error {
	session, err := r.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	session.HasVoted = true
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toSessionKey(sessionID), sessionJSON, 0).Err()
}

func (r *SessionRepository) GetHasVoted(ctx *fiber.Ctx, sessionID string) (bool, error) {
	session, err := r.GetSession(ctx, sessionID)
	if err != nil {
		return false, err
	}

	return session.HasVoted, nil
}
