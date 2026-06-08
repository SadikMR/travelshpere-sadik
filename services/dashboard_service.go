package services

import "github.com/SadikMR/travelshpere-sadik/models"

// DashboardSummary holds the dashboard counter data.
type DashboardSummary struct {
	TotalWishlist int `json:"totalWishlist"`
	PlannedTrips  int `json:"plannedTrips"`
	VisitedTrips  int `json:"visitedTrips"`
}

// GetDashboardSummary returns wishlist counts for a user.
func GetDashboardSummary(username string) DashboardSummary {
	wishlists := wishlistStore.GetByUser(username)

	summary := DashboardSummary{
		TotalWishlist: len(wishlists),
	}

	for _, item := range wishlists {
		switch item.Status {
		case models.StatusPlanned:
			summary.PlannedTrips++
		case models.StatusVisited:
			summary.VisitedTrips++
		}
	}

	return summary
}
