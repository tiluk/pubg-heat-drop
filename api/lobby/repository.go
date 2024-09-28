package lobby

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

func toLobbyKey(lobbyID string) string {
	return "lobby:" + lobbyID
}

func (r *Repository) CreateLobby(ctx *fiber.Ctx, lobby *Lobby) error {
	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toLobbyKey(lobby.LobbyID), lobbyJSON, 0).Err()
}

func (r *Repository) GetLobby(ctx *fiber.Ctx, lobbyID string) (*Lobby, error) {
	lobbyJSON, err := r.cache.Get(ctx.Context(), toLobbyKey(lobbyID)).Result()
	if err != nil {
		return nil, err
	}

	var lobby Lobby
	err = json.Unmarshal([]byte(lobbyJSON), &lobby)
	if err != nil {
		return nil, err
	}

	return &lobby, nil
}
