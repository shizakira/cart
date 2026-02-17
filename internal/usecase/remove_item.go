package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) RemoveItem(ctx context.Context, input dto.RemoveItemInput) error {
	if err := c.storage.RemoveItem(ctx, input.UserID, input.SkuID); err != nil {
		return fmt.Errorf("storage.RemoveItem: %w", err)
	}

	return nil
}
