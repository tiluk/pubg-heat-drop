package session

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

type Session struct {
	SessionID string `json:"sessionID"`
	HasVoted  bool   `json:"hasVoted"`
	Lobby     string `json:"lobby"`
}

func (s *Service) CreateSession(ctx *fiber.Ctx) (*Session, error) {
	session := &Session{
		SessionID: uuid.NewString(),
		HasVoted:  false,
		Lobby:     "",
	}

	err := s.repository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) GetSession(ctx *fiber.Ctx, sessionID string) (*Session, error) {
	session, err := s.repository.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}
