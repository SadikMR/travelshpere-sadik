package models

import "time"

// WishlistStatus represents the travel status of a wishlist entry.
type WishlistStatus string

// Allowed WishlistStatus values.
const (
	StatusPlanned WishlistStatus = "Planned"
	StatusVisited WishlistStatus = "Visited"
)

// Wishlist represents a user's saved travel destination.
type Wishlist struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	CountryName string         `json:"country_name"`
	Note        string         `json:"note"`
	Status      WishlistStatus `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
}