package models

type Attraction struct {
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Category string  `json:"category"`
	Rate     int     `json:"rate"`
	Distance float64 `json:"distance"`
}
