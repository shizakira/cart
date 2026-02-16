package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/model"
)

func (c *Cart) AddItem(ctx context.Context, input dto.AddItemInput) error {
	exist, err := c.productService.IsProductExist(ctx, input.SkuID)
	if err != nil {
		return fmt.Errorf("productService.IsProductExist: %w", err)
	}
	if !exist {
		return model.ErrProductNotFound
	}

	if err = c.storage.AddItem(ctx, input.UserID, domain.Item{
		SkuID: input.SkuID,
		Count: uint(input.Count),
	}); err != nil {
		return fmt.Errorf("storage.AddItem: %w", err)
	}

	return nil
}
