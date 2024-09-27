package vote

import (
	"github.com/tiluk/pubg-heat-drop/lobby"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CastVote(lobbyID string) (*lobby.Heat, error) {

	return nil, nil
}
