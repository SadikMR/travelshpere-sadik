package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	store := NewWishlistStore()

	w := store.Create(&Wishlist{
		Username:    "sadik",
		CountryName: "Bangladesh",
	})

	assert.Equal(t, 1, w.ID)
	assert.Equal(t, "sadik", w.Username)
	assert.Equal(t, "Bangladesh", w.CountryName)
}

func TestCreateAutoIncrements(t *testing.T) {
	store := NewWishlistStore()

	w1 := store.Create(&Wishlist{CountryName: "Japan"})
	w2 := store.Create(&Wishlist{CountryName: "France"})

	assert.Equal(t, 1, w1.ID)
	assert.Equal(t, 2, w2.ID)
}

func TestGetByID(t *testing.T) {
	store := NewWishlistStore()
	created := store.Create(&Wishlist{CountryName: "Italy"})

	found, ok := store.GetByID(created.ID)

	assert.True(t, ok)
	assert.Equal(t, "Italy", found.CountryName)
}

func TestGetByIDNotFound(t *testing.T) {
	store := NewWishlistStore()

	_, ok := store.GetByID(999)

	assert.False(t, ok)
}

func TestGetByUser(t *testing.T) {
	store := NewWishlistStore()
	store.Create(&Wishlist{Username: "sadik", CountryName: "Japan"})
	store.Create(&Wishlist{Username: "other", CountryName: "France"})
	store.Create(&Wishlist{Username: "sadik", CountryName: "Italy"})

	result := store.GetByUser("sadik")

	assert.Len(t, result, 2)
}

func TestGetByUserEmpty(t *testing.T) {
	store := NewWishlistStore()

	result := store.GetByUser("nobody")

	assert.Empty(t, result)
}

func TestUpdate(t *testing.T) {
	store := NewWishlistStore()
	w := store.Create(&Wishlist{CountryName: "Japan", Note: "old"})

	w.Note = "updated"
	store.Update(w)

	found, _ := store.GetByID(w.ID)
	assert.Equal(t, "updated", found.Note)
}

func TestDelete(t *testing.T) {
	store := NewWishlistStore()
	w := store.Create(&Wishlist{CountryName: "Japan"})

	store.Delete(w.ID)

	_, ok := store.GetByID(w.ID)
	assert.False(t, ok)
}
