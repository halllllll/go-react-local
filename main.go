package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"sample/go-react-local-app/config"
	"sample/go-react-local-app/db"
	"sample/go-react-local-app/frontend"
	"sample/go-react-local-app/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/browser"
	sloggin "github.com/samber/slog-gin"

	_ "github.com/mattn/go-sqlite3"
)


func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context)error{
	cfg, err := config.New()
	if err != nil{
		return err
	}
	db, cleanup, err := db.NewDB(ctx, cfg)
	defer cleanup()
	if err != nil {
		return err
	}

	ginlog, err := os.Create(fmt.Sprintf("./%s/gin.log", cfg.Dir))
	if err != nil {
		return err
	}
	defer ginlog.Close()

	applog, err := os.Create(fmt.Sprintf("./%s/app.log", cfg.Dir))

	if err != nil {
		return err
	}
	defer applog.Close()



	appLogger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, applog), nil))

	r := gin.Default()
	r.Use(sloggin.New(slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, ginlog), nil))))
	r.Use(gin.Recovery())
	

	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d", cfg.Port)); err != nil {
		appLogger.Error(err.Error())
	}

	router.SetRoutes(r, db, ctx)
	frontend.RegisterHandlers(r)

	err = r.Run(fmt.Sprintf(":%d", cfg.Port))
	return err
}
