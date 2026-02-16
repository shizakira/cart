package in_memory_storage

import "context"

func (s *Storage) RemoveItem(ctx context.Context, userID int, skuID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cart, ok := s.carts[userID]
	if !ok {
		return nil
	}

	delete(cart, skuID)

	if len(cart) == 0 {
		delete(s.carts, userID)
	}

	return nil
}
