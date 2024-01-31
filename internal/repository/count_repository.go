package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sample/go-react-local-app/internal/models"
)

// repository
type Counter interface {
	Add(context.Context, models.CountValue) error
	FindById(context.Context, models.CountId) (*models.Count, error)
}

type counterRepository struct {
	db *sql.DB
}

func NewCountRepository(db *sql.DB) Counter {
	return &counterRepository{
		db: db,
	}
}

func (cr *counterRepository) Add(ctx context.Context, value models.CountValue) error {
	// トランザクション処理はrollback,commitをtransaction.DoTxでやってるので不要

	stmt := `INSERT INTO count(count_value) VALUES (?)`
	_, err := cr.db.ExecContext(ctx, stmt, value)
	if err != nil {
		return err
	}

	return nil
}

// FindById implements Counter.
func (cr *counterRepository) FindById(ctx context.Context, id models.CountId) (*models.Count, error) {
	// トランザクション処理はrollback,commitをtransaction.DoTxでやってるので不要

	var count models.Count
	stmt := `
		SELECT * FROM count WHERE count_id = ?
	`
	if err := cr.db.QueryRowContext(ctx, stmt, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found")
		} else {
			return nil, err
		}
	}

	return &count, nil
}
