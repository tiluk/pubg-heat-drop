package vote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiluk/pubg-heat-drop/lobby"
	"github.com/tiluk/pubg-heat-drop/session"
)

type Service struct {
	sessionService *session.Service
	lobbyService   *lobby.Service
}

func NewService(sessionService *session.Service, lobbyService *lobby.Service) *Service {
	return &Service{
		sessionService: sessionService,
		lobbyService:   lobbyService,
	}
}

func (s *Service) CastVote(ctx *fiber.Ctx, sessionID string, lobbyID string, heat *lobby.Heat) (*lobby.Lobby, error) {
	lobby, err := s.lobbyService.AddVote(ctx, lobbyID, heat)
	if err != nil {
		return nil, err
	}

	err = s.sessionService.SetVoted(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return lobby, nil
}
