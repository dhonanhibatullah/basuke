package pgxdx

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Transactor struct {
	pool *pgxpool.Pool
}

func NewTransactor(pool *pgxpool.Pool) *Transactor {
	return &Transactor{pool: pool}
}

func (t *Transactor) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	if pgxdx := getTx(ctx); pgxdx != nil {
		return fn(ctx)
	}

	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	wrappedTx := &txWrapper{tx: tx}
	ctx = putTx(ctx, wrappedTx)
	if err := fn(ctx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
