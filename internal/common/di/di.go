package di

import (
	"database/sql"
	"log/slog"
	"sample/go-react-local-app/internal/controller"
	"sample/go-react-local-app/internal/repository"
	"sample/go-react-local-app/internal/service"
)

func InitCount(db *sql.DB, logger *slog.Logger) controller.CountControler {
	tx := repository.NewTransaction(db)
	r := repository.NewCountRepository(db)
	s := service.NewCountSerivce(r, tx, logger)

	return controller.NewCountController(s, logger)
}
