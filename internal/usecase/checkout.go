package usecase

import (
	"context"
	"fmt"

	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) Checkout(ctx context.Context, input dto.CheckoutInput) (dto.CheckoutOutput, error) {
	cart, err := c.storage.Find(ctx, input.UserID)
	if err != nil {
		return dto.CheckoutOutput{}, fmt.Errorf("storage.Find: %w", err)
	}

	orderID, err := c.lomsService.CreateOrder(ctx, input.UserID, cart.Items())
	if err != nil {
		return dto.CheckoutOutput{}, fmt.Errorf("lomsService.CreateOrder: %w", err)
	}

	return dto.CheckoutOutput{OrderID: orderID}, nil
}
