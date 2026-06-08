package services

import (
	"fmt"
	"strings"

	"github.com/SadikMR/travelshpere-sadik/models"
	"github.com/SadikMR/travelshpere-sadik/utils/clients"
	"github.com/SadikMR/travelshpere-sadik/utils/transformers"
)

// CountryService handles country-related operations.
type CountryService struct{}

// GetCountries returns countries filtered by search and region.
func (s *CountryService) GetCountries(
	search string,
	region string,
) ([]models.Country, error) {

	countriesDTO, err := clients.GetCountries()
	if err != nil {
		return nil, err
	}

	var countries []models.Country

	search = strings.ToLower(strings.TrimSpace(search))

	for _, countryDTO := range countriesDTO {
		country := transformers.ToCountry(countryDTO)

		if search != "" &&
			!strings.Contains(
				strings.ToLower(country.Name),
				search,
			) {
			continue
		}

		if region != "" &&
			!strings.EqualFold(country.Region, region) {
			continue
		}

		countries = append(countries, country)
	}

	return countries, nil
}

// GetCountryBySlug returns a country matching the provided slug.
func (s *CountryService) GetCountryBySlug(
	slug string,
) (*models.Country, error) {

	countries, err := s.GetCountries("", "")
	if err != nil {
		return nil, err
	}

	for _, country := range countries {

		countrySlug := strings.ToLower(
			strings.ReplaceAll(country.Name, " ", "-"),
		)

		if countrySlug == slug {
			return &country, nil
		}
	}

	return nil, fmt.Errorf("country not found")
}
