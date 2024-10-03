package lobby

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/tiluk/pubg-heat-drop/models"
)

type LobbyRepository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *LobbyRepository {
	return &LobbyRepository{
		cache: cache,
	}
}

func toLobbyKey(lobbyID string) string {
	return "lobby:" + lobbyID
}

func (r *LobbyRepository) CreateLobby(ctx *fiber.Ctx, lobby *models.Lobby) error {
	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toLobbyKey(lobby.LobbyID), lobbyJSON, 0).Err()
}

func (r *LobbyRepository) GetLobby(ctx *fiber.Ctx, lobbyID string) (*models.Lobby, error) {
	lobbyJSON, err := r.cache.Get(ctx.Context(), toLobbyKey(lobbyID)).Result()
	if err == redis.Nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "lobby not found")
	}
	if err != nil {
		return nil, err
	}

	var lobby models.Lobby
	err = json.Unmarshal([]byte(lobbyJSON), &lobby)
	if err != nil {
		return nil, err
	}

	return &lobby, nil
}

func (r *LobbyRepository) UpdateLobby(ctx *fiber.Ctx, lobby *models.Lobby) error {
	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	return r.cache.Set(ctx.Context(), toLobbyKey(lobby.LobbyID), lobbyJSON, 0).Err()
}

func (r *LobbyRepository) AddVoteToLobby(ctx *fiber.Ctx, lobbyID string, sessionID string, heat models.Heat) error {
	lobby, err := r.GetLobby(ctx, lobbyID)
	if err != nil {
		return err

	}

	err = r.cache.SAdd(ctx.Context(), toLobbyKey(lobbyID)+":sessions", sessionID).Err()
	if err != nil {
		return err
	}

	lobby.Heatmap = append(lobby.Heatmap, heat)

	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return err
	}

	err = r.cache.Set(ctx.Context(), toLobbyKey(lobbyID), lobbyJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *LobbyRepository) GetActiveUsers(ctx *fiber.Ctx, lobbyID string) (int64, error) {
	return r.cache.SCard(ctx.Context(), toLobbyKey(lobbyID)+":sessions").Result()
}
