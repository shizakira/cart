package in_memory_storage

import (
	"context"
	"sync"

	"github.com/shizakira/cart/internal/domain"
)

type Storage struct {
	mx    sync.RWMutex
	carts map[int]*domain.Cart
}

func New() *Storage {
	return &Storage{
		carts: make(map[int]*domain.Cart),
	}
}

func (s *Storage) Find(ctx context.Context, userID int) (*domain.Cart, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.carts[userID], nil
}

func (s *Storage) Save(ctx context.Context, cart *domain.Cart) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.carts[cart.UserID()] = cart

	return nil
}

func (s *Storage) Clear(ctx context.Context, userID int) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.carts[userID]; !ok {
		return nil
	}

	s.carts[userID] = domain.NewCart(userID)

	return nil
}
