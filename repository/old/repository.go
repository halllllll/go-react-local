package old_repository

import (
	"context"
	"database/sql"
)

type Repository struct {
	*sql.DB
	Ctx context.Context
}

func (r *Repository) Get() (int, error) {
	tx, err := r.BeginTx(r.Ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	if err = tx.QueryRowContext(r.Ctx, "SELECT count FROM count WHERE id = $1", 1).Scan(&count); err != nil && err == sql.ErrNoRows {
		// なんもないよの場合
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return count, nil

}

func (r *Repository) Set(newCount int) error {
	tx, err := r.BeginTx(r.Ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var count int
	if err = tx.QueryRowContext(r.Ctx, "SELECT count FROM count WHERE id = $1", 1).Scan(&count); err != nil && err == sql.ErrNoRows {
		_, err := tx.ExecContext(r.Ctx, "INSERT INTO count(count) VALUES ($1)", newCount)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// update
	_, err = tx.ExecContext(r.Ctx, "UPDATE count SET count = $1 WHERE id = $2", newCount, 1)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
