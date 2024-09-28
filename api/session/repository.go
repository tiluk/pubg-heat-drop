package session

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *Repository {
	return &Repository{
		cache: cache,
	}
}

func toSessionKey(sessionID string) string {
	return "session:" + sessionID
}

func (r *Repository) CreateSession(ctx *fiber.Ctx, session *Session) error {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toSessionKey(session.SessionID), sessionJSON, 0).Err()
}

func (r *Repository) GetSession(ctx *fiber.Ctx, sessionID string) (*Session, error) {
	sessionJSON, err := r.cache.Get(ctx.Context(), toSessionKey(sessionID)).Result()
	if err != nil {
		return nil, err
	}

	var session Session
	err = json.Unmarshal([]byte(sessionJSON), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
