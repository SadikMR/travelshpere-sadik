package validators

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCountrySearchValid(t *testing.T) {
	assert.NoError(t, ValidateCountrySearch("bangladesh", ""))
	assert.NoError(t, ValidateCountrySearch("", "Asia"))
	assert.NoError(t, ValidateCountrySearch("japan", "Asia"))
	assert.NoError(t, ValidateCountrySearch("", ""))
}

func TestValidateCountrySearchTooLong(t *testing.T) {
	long := strings.Repeat("a", 101)

	err := ValidateCountrySearch(long, "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "too long")
}

func TestValidateCountrySearchInvalidRegion(t *testing.T) {
	err := ValidateCountrySearch("", "Atlantis")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid region")
}

func TestValidateCountrySearchAllRegions(t *testing.T) {
	regions := []string{"Africa", "Americas", "Asia", "Europe", "Oceania"}

	for _, r := range regions {
		assert.NoError(t, ValidateCountrySearch("", r), "region %s should be valid", r)
	}
}

func TestValidateCountryNameValid(t *testing.T) {
	assert.NoError(t, ValidateCountryName("Bangladesh"))
}

func TestValidateCountryNameEmpty(t *testing.T) {
	assert.Error(t, ValidateCountryName(""))
	assert.Error(t, ValidateCountryName("   "))
}
