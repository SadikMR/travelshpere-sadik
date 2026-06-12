package dto

// RestCountryDTO represents a country response from the REST Countries v5 API.
type RestCountryDTO struct {
	Flag struct {
		Emoji  string `json:"emoji"`
		URLPNG string `json:"url_png"`
		URLSVG string `json:"url_svg"`
	} `json:"flag"`

	Names struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"names"`

	// CHANGED: From map[string]struct to a slice of structs
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`

	// CHANGED: From map[string]string to a slice of structs
	Languages []struct {
		BCP47  string `json:"bcp47"`
		Name   string `json:"name"`
		Native string `json:"native_name"` // Maps to "native_name" in the JSON
	} `json:"languages"`

	// Geography replaces the old 'latlng' array (currently missing from your provided JSON snippet, but keep if you use it later)
	Geography struct {
		Coordinates struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	} `json:"geography"`

	Capitals []struct {
		Name        string `json:"name"`
		Primary     bool   `json:"primary"`
		Coordinates struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	} `json:"capitals"`

	Region string `json:"region"`

	SubRegion string `json:"subregion"`

	Population int64 `json:"population"`
}
