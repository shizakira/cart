package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/shizakira/cart/internal/domain"
)

func (p *Postgres) Find(ctx context.Context, userID int) (*domain.Cart, error) {
	if _, err := p.q.GetCart(ctx, int64(userID)); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("q.GetCart: %w", err)
	}

	rows, err := p.q.GetCartItems(ctx, int64(userID))
	if err != nil {
		return nil, fmt.Errorf("q.GetCartItems: %w", err)
	}

	cart := domain.NewCart(userID)
	for _, r := range rows {
		cart.AddItem(domain.NewItem(int(r.SkuID), uint16(r.Count)))
	}

	return cart, nil
}
