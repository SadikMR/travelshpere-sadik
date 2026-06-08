package models

import "sync"

type WishlistStore struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]*Wishlist
}

func NewWishlistStore() *WishlistStore {
	return &WishlistStore{
		nextID: 1,
		items:  make(map[int]*Wishlist),
	}
}

func (s *WishlistStore) Create(wishlist *Wishlist) *Wishlist {
	s.mu.Lock()
	defer s.mu.Unlock()

	wishlist.ID = s.nextID
	s.nextID++

	s.items[wishlist.ID] = wishlist

	return wishlist
}

func (s *WishlistStore) GetByID(id int) (*Wishlist, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	wishlist, exists := s.items[id]

	return wishlist, exists
}

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

func (s *WishlistStore) Update(wishlist *Wishlist) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[wishlist.ID] = wishlist
}

func (s *WishlistStore) Delete(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.items, id)
}