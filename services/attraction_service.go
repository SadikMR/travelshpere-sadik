package services

import (
	"github.com/SadikMR/travelshpere-sadik/models"
	"github.com/SadikMR/travelshpere-sadik/utils/clients"
	"github.com/SadikMR/travelshpere-sadik/utils/transformers"
)

// GetAttractions returns nearby tourist attractions.
func GetAttractions(
	lat float64,
	lon float64,
) ([]models.Attraction, error) {

	attractionsDTO, err := clients.GetAttractions(lat, lon)
	if err != nil {
		return nil, err
	}

	attractions := transformers.ToAttractions(attractionsDTO)

	return attractions, nil
}
