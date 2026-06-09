package services

import (
	"errors"
	"time"

	"github.com/SadikMR/travelshpere-sadik/models"
)

// Sentinel errors returned by wishlist service operations.
var (
	ErrWishlistNotFound    = errors.New("wishlist not found")
	ErrForbidden           = errors.New("forbidden")
	ErrInvalidStatus       = errors.New("invalid status: allowed values are Planned, Visited")
	ErrCountryNameRequired = errors.New("country_name is required")
)

var wishlistStore = models.NewWishlistStore()

// ListWishlists returns all wishlist entries for a user.
func ListWishlists(username string) []*models.Wishlist {
	return wishlistStore.GetByUser(username)
}

// CreateWishlist adds a country to a user's wishlist.
// Validates required fields. Sets created_at automatically.
func CreateWishlist(
	username,
	countryName,
	note string,
) (*models.Wishlist, error) {
	if countryName == "" {
		return nil, ErrCountryNameRequired
	}

	wishlist := &models.Wishlist{
		Username:    username,
		CountryName: countryName,
		Note:        note,
		Status:      models.StatusPlanned,
		CreatedAt:   time.Now(),
	}

	return wishlistStore.Create(wishlist), nil
}

// UpdateWishlist updates a wishlist entry.
// Validates status values before accepting the request.
func UpdateWishlist(
	username string,
	id int,
	note string,
	status string,
) (*models.Wishlist, error) {
	wishlist, exists := wishlistStore.GetByID(id)
	if !exists {
		return nil, ErrWishlistNotFound
	}

	if wishlist.Username != username {
		return nil, ErrForbidden
	}

	if !isValidStatus(status) {
		return nil, ErrInvalidStatus
	}

	wishlist.Note = note
	wishlist.Status = models.WishlistStatus(status)

	wishlistStore.Update(wishlist)

	return wishlist, nil
}

// DeleteWishlist removes a wishlist entry.
func DeleteWishlist(username string, id int) error {
	wishlist, exists := wishlistStore.GetByID(id)
	if !exists {
		return ErrWishlistNotFound
	}

	if wishlist.Username != username {
		return ErrForbidden
	}

	wishlistStore.Delete(id)

	return nil
}

// isValidStatus checks if the given status is an allowed value.
func isValidStatus(status string) bool {
	switch models.WishlistStatus(status) {
	case models.StatusPlanned, models.StatusVisited:
		return true
	default:
		return false
	}
}
