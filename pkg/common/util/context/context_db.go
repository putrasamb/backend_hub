package helpercontext

import (
	"context"

	"gorm.io/gorm"
)

type TxKeyType struct{}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TxKeyType{}, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TxKeyType{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}
