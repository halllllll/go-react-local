package repository

import (
	"context"
	"database/sql"
	"sample/go-react-local-app/internal/transaction"
)

var txKey = struct{}{}

type tx struct {
	*sql.DB
}

func NewTransaction(db *sql.DB) transaction.Transaction {
	return &tx{db}
}

func (t *tx) DoTx(ctx context.Context, f func(ctx context.Context) error) error {
	txx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer txx.Rollback()
	if err := f(ctx); err != nil {
		return err
	}
	if err := txx.Commit(); err != nil {
		return err
	}
	return nil
}
