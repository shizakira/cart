package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) AddItem(ctx context.Context, input dto.AddItemInput) error {
	exist, err := c.productService.IsProductExist(ctx, input.SkuID)
	if err != nil {
		return fmt.Errorf("productService.IsProductExist: %w", err)
	}
	if !exist {
		return domain.ErrItemNotFound
	}

	if err = c.storage.AddItem(ctx, input.UserID, domain.Item{
		SkuID: input.SkuID,
		Count: input.Count,
	}); err != nil {
		return fmt.Errorf("storage.AddItem: %w", err)
	}

	return nil
}
