package models

// Country represents country information used by the application.
type Country struct {
	Name       string
	Capital    string
	Currency   string
	Languages  []string
	Population int64
	Flag       string
	Region     string
	SubRegion  string
	Latitude   float64
	Longitude  float64
}
