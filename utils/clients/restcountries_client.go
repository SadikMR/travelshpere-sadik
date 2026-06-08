package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/SadikMR/travelshpere-sadik/dto"
)

// GetCountries retrieves all countries.
func GetCountries() ([]dto.RestCountryDTO, error) {
	baseURL, _ := beego.AppConfig.String("restcountriesBaseURL")
	url := fmt.Sprintf(
		"%s/all?fields=name,capital,currencies,languages,flags,latlng,region,subregion,population",
		baseURL,
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

// GetCountryByName retrieves a country by name.
func GetCountryByName(name string) ([]dto.RestCountryDTO, error) {
	baseURL, _ := beego.AppConfig.String("restcountriesBaseURL")
	url := fmt.Sprintf(
		"%s/name/%s?fields=name,capital,currencies,languages,flags,latlng,region,subregion,population",
		baseURL,
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
