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
	tx, err := cr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `INSERT INTO count(value) VALUES (?)`
	_, err = tx.ExecContext(ctx, stmt, value)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// FindById implements Counter.
func (cr *counterRepository) FindById(ctx context.Context, id models.CountId) (*models.Count, error) {
	tx, err := cr.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var count models.Count
	stmt := `
		SELECT * FROM count WHERE id = ?
	`
	if err := tx.QueryRowContext(ctx, stmt, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found")
		} else {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &count, nil
}
