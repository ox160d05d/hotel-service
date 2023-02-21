// Package transaction_manager - Noop transaction manager, нужно заменить на настоящий когда появится транзкционная БД
package transaction_manager

import (
	"context"
	"fmt"
)

type txKey struct{}

type NoopTransactionManager struct {
}

func NewTrxManager() *NoopTransactionManager {
	return &NoopTransactionManager{}
}

func (m *NoopTransactionManager) WithTransaction(
	ctx context.Context,
	txFunc func(ctx context.Context) error,
) (err error) {
	// begin transaction

	defer func() {
		if p := recover(); p != nil {
			// rollback
			err = fmt.Errorf("panic")
		} else if err != nil {
			// rollback
			_ = err
		} else {
			// commit
			_ = err
		}
	}()

	// inject transaction into context
	ctx = context.WithValue(ctx, txKey{}, nil)

	// do
	err = txFunc(ctx)

	return err
}
