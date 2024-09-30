package models

type Lobby struct {
	LobbyID     string `json:"lobbyID"`
	Heatmap     []Heat `json:"heatmap"`
	ActiveUsers int    `json:"activeUsers"`
}

type LobbyResponse struct {
	LobbyID     string    `json:"lobbyID"`
	Heatmap     []Heatmap `json:"heatmap"`
	ActiveUsers int       `json:"activeUsers"`
}
