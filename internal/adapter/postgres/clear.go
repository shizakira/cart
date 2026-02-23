package postgres

import "context"

func (p *Postgres) Clear(ctx context.Context, userID int) error {
	return p.q.DeleteCartItems(ctx, int64(userID))
}
