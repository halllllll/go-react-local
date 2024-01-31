package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"sample/go-react-local-app/internal/models"
)

// repository
type Counter interface {
	Add(context.Context, models.CountValue) error
	FindById(context.Context, models.CountId) (*models.Count, error)
	FindAll(context.Context) (*models.Counts, error)
}

type counterRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewCountRepository(db *sql.DB, logger *slog.Logger) Counter {
	return &counterRepository{
		db:  db,
		log: logger,
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
	// Getほにゃらら系は値を返したいのでtransactionインターフェースを使わず、関数内でトランザクションを貼ることにする
	var count models.Count

	tx, err := cr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	stmt := `
		SELECT * FROM count WHERE count_id = ?
	`
	if err := tx.QueryRowContext(ctx, stmt, id).Scan(&count.Id, &count.Val, &count.Created, &count.Updated); err != nil {
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

// FindAll implements Counter.
func (cr *counterRepository) FindAll(ctx context.Context) (*models.Counts, error) {
	var counts models.Counts
	tx, err := cr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	stmt := `
		SELECT * FROM count
	`
	rows, err := tx.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := &models.Count{}
		if err := rows.Scan(&c.Id, &c.Val, &c.Created, &c.Updated); err != nil {
			cr.log.Error("getall error", err.Error(), "failedId", &c.Id)
		}
		counts = append(counts, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &counts, nil
}
