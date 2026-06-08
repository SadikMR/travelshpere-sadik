package validators

import (
	"errors"
	"strings"
)

const maxSearchLength = 100

var allowedRegions = map[string]struct{}{
	"Africa":   {},
	"Americas": {},
	"Asia":     {},
	"Europe":   {},
	"Oceania":  {},
}

// ValidateCountrySearch validates country search filters.
func ValidateCountrySearch(search, region string) error {
	if len(strings.TrimSpace(search)) > maxSearchLength {
		return errors.New("search query is too long")
	}

	if region == "" {
		return nil
	}

	if _, ok := allowedRegions[region]; !ok {
		return errors.New("invalid region")
	}

	return nil
}

// ValidateCountryName validates a country name.
func ValidateCountryName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("country name is required")
	}

	return nil
}
