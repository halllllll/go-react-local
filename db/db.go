package db

import (
	"context"
	"database/sql"
	"fmt"
	"sample/go-react-local-app/config"
	"time"
)

func NewDB(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/foo.db", cfg.Dir))

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
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				count INTEGER NOT NULL,
				created TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
			)`

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	return db, func() { _ = db.Close() }, nil
}
