package models

type Session struct {
	SessionID string `json:"sessionID"`
	HasVoted  bool   `json:"hasVoted"`
	Lobby     string `json:"lobby"`
}
