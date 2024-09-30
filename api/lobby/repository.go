package lobby

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/tiluk/pubg-heat-drop/models"
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

func (r *Repository) CreateLobby(ctx *fiber.Ctx, lobby *models.Lobby) error {
	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toLobbyKey(lobby.LobbyID), lobbyJSON, 0).Err()
}

func (r *Repository) GetLobby(ctx *fiber.Ctx, lobbyID string) (*models.Lobby, error) {
	lobbyJSON, err := r.cache.Get(ctx.Context(), toLobbyKey(lobbyID)).Result()
	if err == redis.Nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "lobby not found")
	} else if err != nil {
		return nil, err
	}

	var lobby models.Lobby
	err = json.Unmarshal([]byte(lobbyJSON), &lobby)
	if err != nil {
		return nil, err
	}

	return &lobby, nil
}

func (r *Repository) UpdateLobby(ctx *fiber.Ctx, lobby *models.Lobby) error {
	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toLobbyKey(lobby.LobbyID), lobbyJSON, 0).Err()
}
