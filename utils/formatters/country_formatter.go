package formatters

import (
	"fmt"
	"strings"
)

// FormatPopulation converts population into a readable format.
func FormatPopulation(population int64) string {
	switch {
	case population >= 1_000_000_000:
		return fmt.Sprintf("%.1f Billion", float64(population)/1_000_000_000)
	case population >= 1_000_000:
		return fmt.Sprintf("%.1f Million", float64(population)/1_000_000)
	case population >= 1_000:
		return fmt.Sprintf("%.1f Thousand", float64(population)/1_000)
	default:
		return fmt.Sprintf("%d", population)
	}
}

// FormatLanguages joins language names.
func FormatLanguages(languages []string) string {
	return strings.Join(languages, ", ")
}

// FormatCurrency returns a display-friendly currency value.
func FormatCurrency(currency string) string {
	return strings.TrimSpace(currency)
}
