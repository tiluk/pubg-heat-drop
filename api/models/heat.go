package models

type Heat struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Heatmap struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Alt float64 `json:"alt"`
}
