package transformers

import (
	"github.com/SadikMR/travelshpere-sadik/dto"
	"github.com/SadikMR/travelshpere-sadik/models"
)

func ToAttractions(
	response dto.OpenTripMapResponseDTO,
) []models.Attraction {

	attractions := make([]models.Attraction, 0)

	for _, feature := range response.Features {

		if feature.Properties.Name == "" {
			continue
		}

		if len(feature.Geometry.Coordinates) < 2 {
			continue
		}

		attractions = append(attractions, models.Attraction{
			Name:     feature.Properties.Name,
			Lat:      feature.Geometry.Coordinates[1],
			Lon:      feature.Geometry.Coordinates[0],
			Category: feature.Properties.Kinds,
			Rate:     feature.Properties.Rate,
			Distance: feature.Properties.Dist,
		})
	}

	return attractions
}
