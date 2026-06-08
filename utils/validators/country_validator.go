package validators

import (
	"errors"
	"strings"
)

// ValidateCountryName validates a country name.
func ValidateCountryName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("country name is required")
	}

	if len(name) > 100 {
		return errors.New("country name is too long")
	}

	return nil
}
