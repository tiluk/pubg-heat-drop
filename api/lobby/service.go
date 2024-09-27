package lobby

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

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

func (s *Service) CreateLobby(ctx *fiber.Ctx) (*Lobby, error) {
	lobby := &Lobby{
		LobbyID:     uuid.NewString(),
		Heatmap:     []Heat{},
		ActiveUsers: 0,
	}

	err := s.repository.CreateLobby(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}

func (s *Service) GetLobby(ctx *fiber.Ctx, lobbyID string) (*Lobby, error) {
	lobby, err := s.repository.GetLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}
