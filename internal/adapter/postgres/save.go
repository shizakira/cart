package postgres

import (
	"context"
	"fmt"

	cartsqlc "github.com/shizakira/cart/internal/adapter/postgres/cart"
	"github.com/shizakira/cart/internal/domain"
)

func (p *Postgres) Save(ctx context.Context, cart *domain.Cart) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("pool.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := p.q.WithTx(tx)

	if err = qtx.EnsureCart(ctx, int64(cart.UserID())); err != nil {
		return fmt.Errorf("qtx.EnsureCart: %w", err)
	}

	if err = qtx.DeleteCartItems(ctx, int64(cart.UserID())); err != nil {
		return fmt.Errorf("qtx.DeleteCartItems: %w", err)
	}

	for _, item := range cart.Items() {
		arg := cartsqlc.InsertCartItemParams{
			UserID: int64(cart.UserID()),
			SkuID:  int64(item.SkuID()),
			Count:  int16(item.Count()),
		}
		if err = qtx.InsertCartItem(ctx, arg); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
