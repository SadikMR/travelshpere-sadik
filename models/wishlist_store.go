package models

import "sync"

// WishlistStore provides thread-safe in-memory storage for wishlists.
type WishlistStore struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]*Wishlist
}

// NewWishlistStore creates an empty WishlistStore.
func NewWishlistStore() *WishlistStore {
	return &WishlistStore{
		nextID: 1,
		items:  make(map[int]*Wishlist),
	}
}

// Create adds a wishlist entry and assigns an auto-incremented ID.
func (s *WishlistStore) Create(wishlist *Wishlist) *Wishlist {
	s.mu.Lock()
	defer s.mu.Unlock()

	wishlist.ID = s.nextID
	s.nextID++

	s.items[wishlist.ID] = wishlist

	return wishlist
}

// GetByID returns a wishlist entry by its ID.
func (s *WishlistStore) GetByID(id int) (*Wishlist, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	wishlist, exists := s.items[id]

	return wishlist, exists
}

// GetByUser returns all wishlist entries for a given username.
func (s *WishlistStore) GetByUser(username string) []*Wishlist {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*Wishlist

	for _, item := range s.items {
		if item.Username == username {
			result = append(result, item)
		}
	}

	return result
}

// Update replaces a wishlist entry in the store.
func (s *WishlistStore) Update(wishlist *Wishlist) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[wishlist.ID] = wishlist
}

// Delete removes a wishlist entry by ID.
func (s *WishlistStore) Delete(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.items, id)
}