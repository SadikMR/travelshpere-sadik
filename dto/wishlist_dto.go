package dto

type CreateWishlistRequest struct {
	CountryName string `json:"country_name"`
	Note        string `json:"note"`
}

type UpdateWishlistRequest struct {
	Note   string `json:"note"`
	Status string `json:"status"`
}