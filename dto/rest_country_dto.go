package dto

// RestCountryDTO represents a country response from the REST Countries API.
type RestCountryDTO struct {
	Flags struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
	} `json:"flags"`

	Flag string `json:"flag"`

	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	Currencies map[string]struct {
		Name string `json:"name"`
	} `json:"currencies"`

	Languages map[string]string `json:"languages"`

	LatLng []float64 `json:"latlng"`

	Capital []string `json:"capital"`

	Region string `json:"region"`

	SubRegion string `json:"subregion"`

	Population int64 `json:"population"`
}
