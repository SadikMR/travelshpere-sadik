package models

// Country represents country information used by the application.
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
