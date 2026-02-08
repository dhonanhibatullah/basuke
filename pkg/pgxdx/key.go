package pgxdx

import (
	"context"
)

type txKey struct{}

func getTx(ctx context.Context) Pgxdx {
	if pgxdx, ok := ctx.Value(txKey{}).(Pgxdx); ok {
		return pgxdx
	}
	return nil
}

func putTx(ctx context.Context, pgxdx Pgxdx) context.Context {
	return context.WithValue(ctx, txKey{}, pgxdx)
}
