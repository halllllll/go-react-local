package db

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"sample/go-react-local-app/internal/config"
	"time"
)

func NewDB(ctx context.Context, cfg *config.Config, datapath string) (*sql.DB, func(), error) {
	fmt.Println(datapath)
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
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				value INTEGER NOT NULL,
				created TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime')),
				updated TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
			)`

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	return db, func() { _ = db.Close() }, nil
}
