package usecase

import (
	"context"
	"errors"

	"github.com/shizakira/cart/internal/domain"
)

func (c *Cart) getOrCreateCart(ctx context.Context, userID int) (domain.Cart, error) {
	cart, err := c.storage.FindByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			cart, err = c.storage.Create(ctx, userID)
			if err != nil {
				return domain.Cart{}, err
			}
		}
	}

	return cart, nil
}
