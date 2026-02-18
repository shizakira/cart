package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/adapter/postgres/cart"
)

type Postgres struct {
	pool *pgxpool.Pool
	q    *cart.Queries
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{q: cart.New(pool), pool: pool}
}

func (p *Postgres) Close() {
	p.pool.Close()

	log.Info().Msg("postgres: closed")
}
