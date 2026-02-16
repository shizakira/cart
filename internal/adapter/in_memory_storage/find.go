package in_memory_storage

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
)

func (s *Storage) Find(ctx context.Context, userID int) (domain.Cart, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cart, ok := s.carts[userID]
	if !ok {
		return domain.Cart{}, nil
	}

	items := make([]domain.Item, 0, len(cart))
	for _, item := range cart {
		items = append(items, item)
	}

	return domain.Cart{
		UserID: userID,
		Items:  items,
	}, nil
}
