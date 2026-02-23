package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) RemoveItem(ctx context.Context, input dto.RemoveItemInput) error {
	cart, err := c.storage.Find(ctx, input.UserID)
	if err != nil {
		return fmt.Errorf("storage.Find: %w", err)
	}
	if cart == nil {
		return nil
	}

	cart.RemoveItem(input.SkuID)

	if err = c.storage.Save(ctx, cart); err != nil {
		return fmt.Errorf("storage.AddItem: %w", err)
	}

	return nil
}
