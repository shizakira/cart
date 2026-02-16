package in_memory_storage

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
)

func (s *Storage) AddItem(ctx context.Context, userID int, item domain.Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cart, ok := s.carts[userID]
	if !ok {
		cart = make(map[int]domain.Item)
		s.carts[userID] = cart
	}

	if exist, ok := cart[item.SkuID]; ok {
		exist.Count += item.Count
		cart[item.SkuID] = exist

		return nil
	}

	cart[item.SkuID] = item

	return nil
}
