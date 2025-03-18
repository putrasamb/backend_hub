package repository

import (
	"context"
	helpercontext "service-collection/pkg/common/util/context"

	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error)
}

type TransactionRepositoryImplementation struct {
	db *gorm.DB
}

func NewTransactionRepositoryImplementation(db *gorm.DB) *TransactionRepositoryImplementation {
	return &TransactionRepositoryImplementation{
		db: db,
	}
}

func (dc *TransactionRepositoryImplementation) Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error) {
	tx := dc.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	ctxWithTx := helpercontext.SetTx(c, tx)

	result, err := fn(ctxWithTx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return result, nil
}
