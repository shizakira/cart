package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) Clear(ctx context.Context, input dto.ClearCartInput) error {
	if err := c.storage.Clear(ctx, input.UserID); err != nil {
		return fmt.Errorf("storage.Clear: %w", err)
	}

	return nil
}
