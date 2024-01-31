package db

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(ctx context.Context, datapath string) (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite3", filepath.Join(datapath, "data.db"))

	if err != nil {
		return nil, func() {}, err
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	// init
	stmt := `CREATE TABLE IF NOT EXISTS count (
				count_id INTEGER PRIMARY KEY AUTOINCREMENT,
				count_value INTEGER NOT NULL,
				created_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime')),
				updated_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
			)`

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	return db, func() { _ = db.Close() }, nil
}
