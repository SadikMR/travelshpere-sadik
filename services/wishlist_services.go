package services

import (
	"errors"

	"github.com/SadikMR/travelshpere-sadik/models"
)

var (
	ErrWishlistNotFound = errors.New("wishlist not found")
	ErrForbidden        = errors.New("forbidden")
	ErrInvalidStatus    = errors.New("invalid status")
)

var wishlistStore = models.NewWishlistStore()

// ListWishlists returns all wishlist entries for a user.
func ListWishlists(userID string) []*models.Wishlist {
	return wishlistStore.GetByUser(userID)
}

// CreateWishlist adds a country to a user's wishlist.
func CreateWishlist(
	userID,
	countryName,
	note string,
) *models.Wishlist {
	wishlist := &models.Wishlist{
		UserID:      userID,
		CountryName: countryName,
		Note:        note,
		Status:      models.StatusWantToVisit,
	}

	return wishlistStore.Create(wishlist)
}

// UpdateWishlist updates a wishlist entry.
func UpdateWishlist(
	userID string,
	id int,
	note string,
	status string,
) (*models.Wishlist, error) {
	wishlist, exists := wishlistStore.GetByID(id)
	if !exists {
		return nil, ErrWishlistNotFound
	}

	if wishlist.UserID != userID {
		return nil, ErrForbidden
	}

	switch status {
	case string(models.StatusWantToVisit),
		string(models.StatusPlanned),
		string(models.StatusVisited):
	default:
		return nil, ErrInvalidStatus
	}

	wishlist.Note = note
	wishlist.Status = models.WishlistStatus(status)

	wishlistStore.Update(wishlist)

	return wishlist, nil
}

// DeleteWishlist removes a wishlist entry.
func DeleteWishlist(userID string, id int) error {
	wishlist, exists := wishlistStore.GetByID(id)
	if !exists {
		return ErrWishlistNotFound
	}

	if wishlist.UserID != userID {
		return ErrForbidden
	}

	wishlistStore.Delete(id)

	return nil
}