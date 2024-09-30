package lobby

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiluk/pubg-heat-drop/models"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateLobby(ctx *fiber.Ctx) (*models.Lobby, error) {
	lobby := &models.Lobby{
		LobbyID:     uuid.NewString(),
		Heatmap:     []models.Heat{},
		ActiveUsers: 0,
	}

	err := s.repository.CreateLobby(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}

func (s *Service) GetLobby(ctx *fiber.Ctx, lobbyID string) (*models.LobbyResponse, error) {
	var intensityMultiplier float64 = 10

	lobby, err := s.repository.GetLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}

	lobbyResponse := &models.LobbyResponse{
		LobbyID:     lobby.LobbyID,
		Heatmap:     []models.Heatmap{},
		ActiveUsers: lobby.ActiveUsers,
	}
	lobbyResponse.Heatmap = make([]models.Heatmap, len(lobby.Heatmap))
	for i, heat := range lobby.Heatmap {
		lobbyResponse.Heatmap[i] = models.Heatmap{
			Lat: heat.Lat,
			Lng: heat.Lng,
			Alt: intensityMultiplier / float64(lobby.ActiveUsers),
		}
	}

	return lobbyResponse, nil
}

func (s *Service) AddVote(ctx *fiber.Ctx, lobbyID string, heat *models.Heat) (*models.Lobby, error) {
	lobby, err := s.repository.GetLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}

	lobby.Heatmap = append(lobby.Heatmap, *heat)
	lobby.ActiveUsers++

	err = s.repository.UpdateLobby(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}
