package dto

type RestCountryDTO struct {
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

	Population int64 `json:"population"`
}
