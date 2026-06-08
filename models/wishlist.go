package models

import "time"

type WishlistStatus string

const (
	StatusPlanned WishlistStatus = "Planned"
	StatusVisited WishlistStatus = "Visited"
)

type Wishlist struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	CountryName string         `json:"country_name"`
	Note        string         `json:"note"`
	Status      WishlistStatus `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
}