package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SadikMR/travelshpere-sadik/dto"
)

const restCountriesURL = "https://restcountries.com/v3.1/all?fields=name,capital,currencies,languages,flag,latlng,region,population"

// GetCountries retrieves country data from the REST Countries API.
func GetCountries() ([]dto.RestCountryDTO, error) {
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var countries []dto.RestCountryDTO

	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	return countries, nil
}

// GetCountryByName retrieves a country by name.
func GetCountryByName(name string) ([]dto.RestCountryDTO, error) {
	url := fmt.Sprintf(
		"https://restcountries.com/v3.1/name/%s?fields=name,capital,currencies,languages,flag,latlng,region,population",
		name,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var countries []dto.RestCountryDTO

	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	return countries, nil
}
