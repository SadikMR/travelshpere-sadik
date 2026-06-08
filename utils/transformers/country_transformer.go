package transformers

import (
	"github.com/SadikMR/travelshpere-sadik/dto"
	"github.com/SadikMR/travelshpere-sadik/models"
)

// ToCountry converts a REST Countries DTO into a Country model.
func ToCountry(countryDTO dto.RestCountryDTO) models.Country {
	country := models.Country{
		Name:       countryDTO.Name.Common,
		Region:     countryDTO.Region,
		Population: countryDTO.Population,
		Flag:       countryDTO.Flags.PNG,
	}

	if len(countryDTO.Capital) > 0 {
		country.Capital = countryDTO.Capital[0]
	}

	if len(countryDTO.LatLng) >= 2 {
		country.Latitude = countryDTO.LatLng[0]
		country.Longitude = countryDTO.LatLng[1]
	}

	for _, currency := range countryDTO.Currencies {
		country.Currency = currency.Name
		break
	}

	for _, language := range countryDTO.Languages {
		country.Languages = append(
			country.Languages,
			language,
		)
	}

	return country
}
