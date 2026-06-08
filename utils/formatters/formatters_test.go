package formatters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatPopulationBillions(t *testing.T) {
	assert.Equal(t, "1.4 Billion", FormatPopulation(1_400_000_000))
}

func TestFormatPopulationMillions(t *testing.T) {
	assert.Equal(t, "170.0 Million", FormatPopulation(170_000_000))
}

func TestFormatPopulationThousands(t *testing.T) {
	assert.Equal(t, "45.0 Thousand", FormatPopulation(45_000))
}

func TestFormatPopulationSmall(t *testing.T) {
	assert.Equal(t, "500", FormatPopulation(500))
	assert.Equal(t, "0", FormatPopulation(0))
}

func TestFormatLanguages(t *testing.T) {
	assert.Equal(t, "Bengali, English", FormatLanguages([]string{"Bengali", "English"}))
	assert.Equal(t, "Japanese", FormatLanguages([]string{"Japanese"}))
	assert.Equal(t, "", FormatLanguages([]string{}))
}

func TestFormatCurrency(t *testing.T) {
	assert.Equal(t, "Taka", FormatCurrency("Taka"))
	assert.Equal(t, "Yen", FormatCurrency(" Yen "))
	assert.Equal(t, "", FormatCurrency(""))
}

func TestFormatDateTime(t *testing.T) {
	date := time.Date(2026, 6, 9, 14, 30, 0, 0, time.UTC)

	assert.Equal(t, "09 Jun 2026 14:30", FormatDateTime(date))
}
