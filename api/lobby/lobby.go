package lobby

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Heat struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Alt float64 `json:"alt"`
}

type Lobby struct {
	LobbyID     string `json:"lobbyID"`
	Heatmap     []Heat `json:"heatmap"`
	ActiveUsers int    `json:"activeUsers"`
}

type LobbyHandler struct {
	cache *redis.Client
}

func NewLobbyHandler(cache *redis.Client) *LobbyHandler {
	return &LobbyHandler{
		cache: cache,
	}
}

func (lh *LobbyHandler) CreateLobby(ctx *fiber.Ctx) error {

	lobby := Lobby{
		LobbyID:     uuid.NewString(),
		Heatmap:     []Heat{},
		ActiveUsers: 0,
	}

	lobbyJSON, err := json.Marshal(lobby)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = lh.cache.Set(ctx.Context(), "lobby:"+lobby.LobbyID, lobbyJSON, 0).Err()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	log.Println("Lobby created with ID:", lobby.LobbyID)

	return ctx.JSON(lobby)
}

func (lh *LobbyHandler) GetLobby(ctx *fiber.Ctx) error {
	lobbyID := ctx.Params("id")
	if lobbyID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing lobby ID")
	}

	log.Println("Getting lobby with ID:", lobbyID)

	lobbyJSON, err := lh.cache.Get(ctx.Context(), "session:"+lobbyID).Result()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var lobby Lobby
	err = json.Unmarshal([]byte(lobbyJSON), &lobby)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(lobby)
}
