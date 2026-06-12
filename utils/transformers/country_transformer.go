package transformers

import (
	"github.com/SadikMR/travelshpere-sadik/dto"
	"github.com/SadikMR/travelshpere-sadik/models"
)

// ToCountry converts a REST Countries v5 DTO into a Country model.
func ToCountry(countryDTO dto.RestCountryDTO) models.Country {
	country := models.Country{
		Name:       countryDTO.Names.Common,
		Region:     countryDTO.Region,
		SubRegion:  countryDTO.SubRegion,
		Population: countryDTO.Population,
		Flag:       countryDTO.Flag.URLPNG,
	}

	if len(countryDTO.Capitals) > 0 {
		country.Capital = countryDTO.Capitals[0].Name
	}

	// Set coordinates from geography data; if unavailable, fall back to capital coordinates.
	if countryDTO.Geography.Coordinates.Lat != 0 || countryDTO.Geography.Coordinates.Lng != 0 {
		country.Latitude = countryDTO.Geography.Coordinates.Lat
		country.Longitude = countryDTO.Geography.Coordinates.Lng
	} else if len(countryDTO.Capitals) > 0 {
		country.Latitude = countryDTO.Capitals[0].Coordinates.Lat
		country.Longitude = countryDTO.Capitals[0].Coordinates.Lng
	}

	// CHANGED: Grab the first currency name from the array
	if len(countryDTO.Currencies) > 0 {
		country.Currency = countryDTO.Currencies[0].Name
	}

	// CHANGED: Loop over the languages array and append the string name
	for _, language := range countryDTO.Languages {
		country.Languages = append(
			country.Languages,
			language.Name,
		)
	}

	return country
}
