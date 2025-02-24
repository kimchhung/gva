package dbtx

import (
	"context"

	"gorm.io/gorm"
)

// contextTxKey is a context key for the transaction.
type contextTxKey struct{}

func WrapTxCtx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, contextTxKey{}, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if !ok {
		return nil
	}
	return tx
}
