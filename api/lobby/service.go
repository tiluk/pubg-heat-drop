package lobby

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tiluk/pubg-heat-drop/models"
	"github.com/tiluk/pubg-heat-drop/session"
)

type LobbyService struct {
	repository     *LobbyRepository
	sessionService *session.SessionService
}

func NewService(repository *LobbyRepository, sessionService *session.SessionService) *LobbyService {
	return &LobbyService{
		repository:     repository,
		sessionService: sessionService,
	}
}

func (s *LobbyService) LobbyToLobbyResponse(ctx *fiber.Ctx, lobby *models.Lobby) (*models.LobbyResponse, error) {
	var intensityMultiplier float64 = 10

	activeUsers, err := s.repository.GetActiveUsers(ctx, lobby.LobbyID)
	if err != nil {
		return nil, err
	}

	lobbyResponse := &models.LobbyResponse{
		LobbyID: lobby.LobbyID,
		Heatmap: []models.Heatmap{},
	}

	lobbyResponse.Heatmap = make([]models.Heatmap, len(lobby.Heatmap))
	for i, heat := range lobby.Heatmap {
		lobbyResponse.Heatmap[i] = models.Heatmap{
			Lat: heat.Lat,
			Lng: heat.Lng,
			Alt: intensityMultiplier / float64(activeUsers),
		}
	}

	return lobbyResponse, nil
}

func (s *LobbyService) CreateLobby(ctx *fiber.Ctx) (*models.Lobby, error) {
	lobby := &models.Lobby{
		LobbyID: uuid.NewString(),
		Heatmap: []models.Heat{},
	}

	err := s.repository.CreateLobby(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}

func (s *LobbyService) GetLobby(ctx *fiber.Ctx, lobbyID string) (*models.LobbyResponse, error) {
	lobby, err := s.repository.GetLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}

	lobbyResponse, err := s.LobbyToLobbyResponse(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobbyResponse, nil
}

func (s *LobbyService) AddVote(ctx *fiber.Ctx, lobbyID string, heat *models.Heat) (*models.Lobby, error) {
	lobby, err := s.repository.GetLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}

	lobby.Heatmap = append(lobby.Heatmap, *heat)

	err = s.repository.UpdateLobby(ctx, lobby)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}

func (s *LobbyService) AddLobbyVote(ctx *fiber.Ctx, lobbyID string, sessionID string, heat models.Heat) error {
	err := s.repository.AddVoteToLobby(ctx, lobbyID, sessionID, heat)
	if err != nil {
		return err
	}

	if ctx.Locals("hasVoted").(bool) {
		return fiber.NewError(fiber.StatusConflict, "User has already voted")
	}
	err = s.sessionService.SetVoted(ctx, sessionID)
	if err != nil {
		return err
	}

	return nil
}
