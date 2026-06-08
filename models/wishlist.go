package models

type WishlistStatus string

const (
	StatusWantToVisit WishlistStatus = "want_to_visit"
	StatusPlanned     WishlistStatus = "planned"
	StatusVisited     WishlistStatus = "visited"
)

type Wishlist struct {
    ID          int            `json:"id"`
    Username    string         `json:"username"`
    CountryName string         `json:"country_name"`
    Note        string         `json:"note"`
    Status      WishlistStatus `json:"status"`
}