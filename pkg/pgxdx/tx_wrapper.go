package pgxdx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type txWrapper struct {
	tx pgx.Tx
}

func (t *txWrapper) Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error) {
	return t.tx.Exec(ctx, sql, args...)
}

func (t *txWrapper) Query(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error) {
	return t.tx.Query(ctx, sql, args...)
}

func (t *txWrapper) QueryRow(ctx context.Context, sql string, args ...any) (row pgx.Row) {
	return t.tx.QueryRow(ctx, sql, args...)
}

func (t *txWrapper) WithTx(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}
