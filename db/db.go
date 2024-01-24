package db

import (
	"database/sql"
	"fmt"
	"sample/go-react-local-app/config"
)

func NewDB(cfg *config.Config)(*sql.DB, error){
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/foo.db", cfg.Dir))
	if err != nil {
		return nil, err		
	}
	return db, nil
}