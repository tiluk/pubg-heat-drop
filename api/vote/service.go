package vote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiluk/pubg-heat-drop/lobby"
	"github.com/tiluk/pubg-heat-drop/models"
	"github.com/tiluk/pubg-heat-drop/session"
)

type VoteService struct {
	sessionService *session.SessionService
	lobbyService   *lobby.LobbyService
}

func NewService(sessionService *session.SessionService, lobbyService *lobby.LobbyService) *VoteService {
	return &VoteService{
		sessionService: sessionService,
		lobbyService:   lobbyService,
	}
}

func (s *VoteService) CastVote(ctx *fiber.Ctx, sessionID string, lobbyID string, heat *models.Heat) (*models.LobbyResponse, error) {
	lobbyWithHeat, err := s.lobbyService.AddVote(ctx, lobbyID, heat)
	if err != nil {
		return nil, err
	}

	err = s.sessionService.SetVoted(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	lobbyResponse := lobby.LobbyToLobbyResponse(lobbyWithHeat)
	return lobbyResponse, nil
}
