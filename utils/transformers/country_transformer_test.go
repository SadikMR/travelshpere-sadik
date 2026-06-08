package transformers

import (
	"testing"

	"github.com/SadikMR/travelshpere-sadik/dto"
	"github.com/stretchr/testify/assert"
)

func TestToCountryBasicFields(t *testing.T) {
	input := dto.RestCountryDTO{
		Name:       struct{ Common string `json:"common"` }{Common: "Bangladesh"},
		Region:     "Asia",
		SubRegion:  "Southern Asia",
		Population: 170_000_000,
		Capital:    []string{"Dhaka"},
		LatLng:     []float64{23.685, 90.356},
	}
	input.Flags.PNG = "https://example.com/flag.png"
	input.Currencies = map[string]struct {
		Name string `json:"name"`
	}{
		"BDT": {Name: "Taka"},
	}
	input.Languages = map[string]string{
		"ben": "Bengali",
	}

	country := ToCountry(input)

	assert.Equal(t, "Bangladesh", country.Name)
	assert.Equal(t, "Asia", country.Region)
	assert.Equal(t, "Southern Asia", country.SubRegion)
	assert.Equal(t, int64(170_000_000), country.Population)
	assert.Equal(t, "Dhaka", country.Capital)
	assert.Equal(t, 23.685, country.Latitude)
	assert.Equal(t, 90.356, country.Longitude)
	assert.Equal(t, "https://example.com/flag.png", country.Flag)
	assert.Equal(t, "Taka", country.Currency)
	assert.Contains(t, country.Languages, "Bengali")
}

func TestToCountryEmptyCapital(t *testing.T) {
	input := dto.RestCountryDTO{
		Capital: []string{},
	}

	country := ToCountry(input)

	assert.Equal(t, "", country.Capital)
}

func TestToCountryNoLatLng(t *testing.T) {
	input := dto.RestCountryDTO{
		LatLng: []float64{},
	}

	country := ToCountry(input)

	assert.Equal(t, 0.0, country.Latitude)
	assert.Equal(t, 0.0, country.Longitude)
}

func TestToCountryNoCurrencies(t *testing.T) {
	input := dto.RestCountryDTO{}

	country := ToCountry(input)

	assert.Equal(t, "", country.Currency)
}

func TestToCountryNoLanguages(t *testing.T) {
	input := dto.RestCountryDTO{}

	country := ToCountry(input)

	assert.Empty(t, country.Languages)
}

func TestToAttractionsValid(t *testing.T) {
	input := dto.OpenTripMapResponseDTO{
		Features: []dto.OpenTripFeatureDTO{
			{
				Properties: dto.OpenTripPropertyDTO{
					Name:  "Lalbagh Fort",
					Kinds: "historic",
					Rate:  7,
					Dist:  1200.5,
				},
				Geometry: dto.OpenTripGeometryDTO{
					Coordinates: []float64{90.388, 23.719},
				},
			},
		},
	}

	result := ToAttractions(input)

	assert.Len(t, result, 1)
	assert.Equal(t, "Lalbagh Fort", result[0].Name)
	assert.Equal(t, "historic", result[0].Category)
	assert.Equal(t, 7, result[0].Rate)
	assert.Equal(t, 1200.5, result[0].Distance)
	assert.Equal(t, 23.719, result[0].Lat)
	assert.Equal(t, 90.388, result[0].Lon)
}

func TestToAttractionsSkipsUnnamed(t *testing.T) {
	input := dto.OpenTripMapResponseDTO{
		Features: []dto.OpenTripFeatureDTO{
			{
				Properties: dto.OpenTripPropertyDTO{Name: ""},
				Geometry:   dto.OpenTripGeometryDTO{Coordinates: []float64{1, 2}},
			},
		},
	}

	result := ToAttractions(input)

	assert.Empty(t, result)
}

func TestToAttractionsSkipsMissingCoords(t *testing.T) {
	input := dto.OpenTripMapResponseDTO{
		Features: []dto.OpenTripFeatureDTO{
			{
				Properties: dto.OpenTripPropertyDTO{Name: "Place"},
				Geometry:   dto.OpenTripGeometryDTO{Coordinates: []float64{1}},
			},
		},
	}

	result := ToAttractions(input)

	assert.Empty(t, result)
}

func TestToAttractionsEmpty(t *testing.T) {
	input := dto.OpenTripMapResponseDTO{}

	result := ToAttractions(input)

	assert.Empty(t, result)
}
