package models

type WishlistStatus string

const (
	StatusWantToVisit WishlistStatus = "want_to_visit"
	StatusPlanned     WishlistStatus = "planned"
	StatusVisited     WishlistStatus = "visited"
)

type Wishlist struct {
    ID          int            `json:"id"`
    UserID      string         `json:"user_id"`
    CountryName string         `json:"country_name"`
    Note        string         `json:"note"`
    Status      WishlistStatus `json:"status"`
}