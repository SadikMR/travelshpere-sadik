package formatters

import "time"

// FormatDateTime formats a date and time for display.
func FormatDateTime(t time.Time) string {
	return t.Format("02 Jan 2006 15:04")
}
