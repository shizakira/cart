package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) AddItem(ctx context.Context, input dto.AddItemInput) error {
	cart, err := c.getOrCreateCart(ctx, input.UserID)
	if err != nil {
		return fmt.Errorf("getOrCreateCart: %w", err)
	}

	exist, err := c.productService.IsProductExist(ctx, input.SkuID)
	if err != nil {
		return fmt.Errorf("productService.IsProductExist: %w", err)
	}
	if !exist {
		return domain.ErrNotFound
	}

	for i := range cart.Items {
		if cart.Items[i].SkuID != input.SkuID {
			continue
		}

		cart.Items[i].Count += input.Count
		if err = c.storage.Save(ctx, cart); err != nil {
			return fmt.Errorf("storage.Save: %w", err)
		}

		return nil
	}

	cart.Items = append(cart.Items, domain.Item{
		SkuID: input.SkuID,
		Count: input.Count,
	})
	if err = c.storage.Save(ctx, cart); err != nil {
		return fmt.Errorf("storage.Save: %w", err)
	}

	return nil
}
