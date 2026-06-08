package services

import (
	"testing"

	"github.com/SadikMR/travelshpere-sadik/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// resetStore replaces the package-level store with a fresh one before each test.
func resetStore() {
	wishlistStore = models.NewWishlistStore()
}

func TestListWishlistsEmpty(t *testing.T) {
	resetStore()

	result := ListWishlists("sadik")

	assert.Empty(t, result)
}

func TestCreateWishlistSuccess(t *testing.T) {
	resetStore()

	w, err := CreateWishlist("sadik", "Bangladesh", "my home")

	require.NoError(t, err)
	assert.Equal(t, "sadik", w.Username)
	assert.Equal(t, "Bangladesh", w.CountryName)
	assert.Equal(t, "my home", w.Note)
	assert.Equal(t, models.StatusPlanned, w.Status)
	assert.False(t, w.CreatedAt.IsZero())
}

func TestCreateWishlistMissingCountry(t *testing.T) {
	resetStore()

	w, err := CreateWishlist("sadik", "", "note")

	assert.Nil(t, w)
	assert.ErrorIs(t, err, ErrCountryNameRequired)
}

func TestListWishlistsReturnsUserEntries(t *testing.T) {
	resetStore()

	CreateWishlist("sadik", "Japan", "")
	CreateWishlist("other", "France", "")
	CreateWishlist("sadik", "Italy", "")

	result := ListWishlists("sadik")

	assert.Len(t, result, 2)
}

func TestUpdateWishlistSuccess(t *testing.T) {
	resetStore()

	w, _ := CreateWishlist("sadik", "Japan", "old note")

	updated, err := UpdateWishlist("sadik", w.ID, "new note", "Visited")

	require.NoError(t, err)
	assert.Equal(t, "new note", updated.Note)
	assert.Equal(t, models.StatusVisited, updated.Status)
}

func TestUpdateWishlistNotFound(t *testing.T) {
	resetStore()

	_, err := UpdateWishlist("sadik", 999, "note", "Planned")

	assert.ErrorIs(t, err, ErrWishlistNotFound)
}

func TestUpdateWishlistForbidden(t *testing.T) {
	resetStore()

	w, _ := CreateWishlist("sadik", "Japan", "")

	_, err := UpdateWishlist("hacker", w.ID, "note", "Planned")

	assert.ErrorIs(t, err, ErrForbidden)
}

func TestUpdateWishlistInvalidStatus(t *testing.T) {
	resetStore()

	w, _ := CreateWishlist("sadik", "Japan", "")

	_, err := UpdateWishlist("sadik", w.ID, "note", "InvalidStatus")

	assert.ErrorIs(t, err, ErrInvalidStatus)
}

func TestDeleteWishlistSuccess(t *testing.T) {
	resetStore()

	w, _ := CreateWishlist("sadik", "Japan", "")

	err := DeleteWishlist("sadik", w.ID)

	assert.NoError(t, err)
	assert.Empty(t, ListWishlists("sadik"))
}

func TestDeleteWishlistNotFound(t *testing.T) {
	resetStore()

	err := DeleteWishlist("sadik", 999)

	assert.ErrorIs(t, err, ErrWishlistNotFound)
}

func TestDeleteWishlistForbidden(t *testing.T) {
	resetStore()

	w, _ := CreateWishlist("sadik", "Japan", "")

	err := DeleteWishlist("hacker", w.ID)

	assert.ErrorIs(t, err, ErrForbidden)
}
