package models

// Country represents a country with geographic and demographic data.
type Country struct {
	Name       string   `json:"name"`
	Capital    string   `json:"capital"`
	Currency   string   `json:"currency"`
	Languages  []string `json:"languages"`
	Population int64    `json:"population"`
	Flag       string   `json:"flag"`
	Region     string   `json:"region"`
	SubRegion  string   `json:"subRegion"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
}
