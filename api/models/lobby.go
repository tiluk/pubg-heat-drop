package models

type Lobby struct {
	LobbyID string `json:"lobbyID"`
	Heatmap []Heat `json:"heatmap"`
}

type LobbyResponse struct {
	LobbyID string    `json:"lobbyID"`
	Heatmap []Heatmap `json:"heatmap"`
}
