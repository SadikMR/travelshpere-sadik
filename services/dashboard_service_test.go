package services

import (
	"testing"

	"github.com/SadikMR/travelshpere-sadik/models"
	"github.com/stretchr/testify/assert"
)

func TestDashboardSummaryEmpty(t *testing.T) {
	resetStore()

	summary := GetDashboardSummary("sadik")

	assert.Equal(t, 0, summary.TotalWishlist)
	assert.Equal(t, 0, summary.PlannedTrips)
	assert.Equal(t, 0, summary.VisitedTrips)
}

func TestDashboardSummaryCounts(t *testing.T) {
	resetStore()

	// Create 3 entries: 2 Planned (default), then update 1 to Visited
	CreateWishlist("sadik", "Japan", "")
	CreateWishlist("sadik", "France", "")
	w3, _ := CreateWishlist("sadik", "Italy", "")

	wishlistStore.GetByID(w3.ID)
	w3.Status = models.StatusVisited
	wishlistStore.Update(w3)

	summary := GetDashboardSummary("sadik")

	assert.Equal(t, 3, summary.TotalWishlist)
	assert.Equal(t, 2, summary.PlannedTrips)
	assert.Equal(t, 1, summary.VisitedTrips)
}

func TestDashboardSummaryOnlyCountsOwnEntries(t *testing.T) {
	resetStore()

	CreateWishlist("sadik", "Japan", "")
	CreateWishlist("other", "France", "")

	summary := GetDashboardSummary("sadik")

	assert.Equal(t, 1, summary.TotalWishlist)
}
