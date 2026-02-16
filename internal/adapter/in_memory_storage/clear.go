package in_memory_storage

import "context"

func (s *Storage) Clear(ctx context.Context, userID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.carts, userID)
	return nil
}
