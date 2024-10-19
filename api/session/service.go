package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tiluk/pubg-heat-drop/models"
)

type SessionService struct {
	repository *SessionRepository
}

func NewService(repository *SessionRepository) *SessionService {
	return &SessionService{
		repository: repository,
	}
}

func (s *SessionService) CreateSession(ctx *fiber.Ctx) (*string, error) {
	session := &models.Session{
		SessionID: uuid.NewString(),
		HasVoted:  false,
	}

	err := s.repository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionID": session.SessionID,
	})

	jwt, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &jwt, nil
}

func (s *SessionService) SetVoted(ctx *fiber.Ctx, sessionID string) error {
	err := s.repository.SetHasVoted(ctx, sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionService) GetHasVoted(ctx *fiber.Ctx, sessionID string) (bool, error) {
	hasVoted, err := s.repository.GetHasVoted(ctx, sessionID)
	if err == redis.Nil {
		return false, fiber.NewError(fiber.StatusNotFound, "session not found")
	}
	if err != nil {
		return false, err
	}

	return hasVoted, nil
}

func (s *SessionService) VerifyJWTSession(ctx *fiber.Ctx, unsafeSession *models.Session) (bool, error) {
	_, err := s.repository.GetSession(ctx, unsafeSession.SessionID)
	if err != nil && err != redis.Nil {
		return false, err
	}
	if err == redis.Nil {
		return false, nil
	}

	return true, nil
}
