package pgxdx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pgxdx interface {
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error)
	QueryRow(ctx context.Context, sql string, args ...any) (row pgx.Row)
	WithTx(ctx context.Context, fn func(context.Context) error) error
}

type dx struct {
	pool *pgxpool.Pool
}

func NewPgxdx(pool *pgxpool.Pool) Pgxdx {
	return &dx{
		pool: pool,
	}
}

func (p *dx) Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error) {
	if tx := getTx(ctx); tx != nil {
		return tx.Exec(ctx, sql, args...)
	}
	return p.pool.Exec(ctx, sql, args...)
}

func (p *dx) Query(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error) {
	if tx := getTx(ctx); tx != nil {
		return tx.Query(ctx, sql, args...)
	}
	return p.pool.Query(ctx, sql, args...)
}

func (p *dx) QueryRow(ctx context.Context, sql string, args ...any) (row pgx.Row) {
	if tx := getTx(ctx); tx != nil {
		return tx.QueryRow(ctx, sql, args...)
	}
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *dx) WithTx(ctx context.Context, fn func(context.Context) error) error {
	if tx := getTx(ctx); tx != nil {
		return tx.WithTx(ctx, fn)
	}

	tx, err := p.pool.Begin(ctx)
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
