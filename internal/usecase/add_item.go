package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) AddItem(ctx context.Context, input dto.AddItemInput) error {
	_, err := c.productService.GetProduct(ctx, input.SkuID)
	if err != nil {
		return fmt.Errorf("productService.GetProduct: %w", err)
	}

	count, err := c.lomsService.StocksInfo(ctx, uint32(input.SkuID))
	if err != nil {
		return fmt.Errorf("lomsService.StocksInfo: %w", err)
	}

	if count < uint64(input.Count) {
		return domain.ErrInsufficientStock
	}

	cart, err := c.storage.Find(ctx, input.UserID)
	if err != nil {
		return fmt.Errorf("storage.Find: %w", err)
	}

	if cart == nil {
		cart = domain.NewCart(input.UserID)
	}

	item := domain.NewItem(input.SkuID, uint16(input.Count))
	cart.AddItem(item)

	if err = c.storage.Save(ctx, cart); err != nil {
		return fmt.Errorf("storage.Save: %w", err)
	}

	return nil
}
