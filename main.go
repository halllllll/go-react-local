package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sample/go-react-local-app/frontend"
	"sample/go-react-local-app/internal/config"
	"sample/go-react-local-app/internal/controller"
	"sample/go-react-local-app/internal/db"
	"sample/go-react-local-app/internal/repository"
	"sample/go-react-local-app/internal/router"
	"sample/go-react-local-app/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/browser"
	sloggin "github.com/samber/slog-gin"

	_ "github.com/mattn/go-sqlite3"
)

var AppMode string

type AppEnv string

var (
	ProdEnv AppEnv = "prod"
	ProdDev AppEnv = "dev"
)

func main() {
	// go run(make dev)時と実行ファイルの実行時でカレントディレクトリを切り替える
	if err := os.Setenv("ENV", AppMode); err != nil {
		log.Fatal(err)
	}

	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	datapath, err := checkEnv(cfg)
	if err != nil {
		return err
	}

	db, cleanup, err := db.NewDB(ctx, cfg, datapath)
	defer cleanup()
	if err != nil {
		return err
	}

	ctrl := controller.NewCountController(
		service.NewCountSerivce(
			repository.NewCountRepository(db),
		),
	)

	ginlog, err := os.Create(filepath.Join(datapath, "gin.log"))
	if err != nil {
		return err
	}
	defer ginlog.Close()

	applog, err := os.Create(filepath.Join(datapath, "app.log"))

	if err != nil {
		return err
	}
	defer applog.Close()

	appLogger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, applog), nil))

	r := gin.Default()
	// middlewares
	r.Use(sloggin.New(slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, ginlog), nil))))
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			fmt.Sprintf("http://%s", cfg.Address),
			fmt.Sprintf("http://127.0.0.1:%d", cfg.Port),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS", // for preflight request
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,           // need cookie
		MaxAge:           24 * time.Hour, //
	}))

	if cfg.Env == string(ProdEnv) {
		frontend.RegisterHandlers(r)
		if err := browser.OpenURL(fmt.Sprintf("http://%s", cfg.Address)); err != nil {
			appLogger.Error(err.Error())
		}
	} else if cfg.Env == string(ProdDev) {
		frontend.SetupProxy(r)
	}

	router.SetRoutes(r, ctrl)
	err = r.Run(fmt.Sprintf(":%d", cfg.Port))
	return err
}

func checkEnv(cfg *config.Config) (string, error) {
	var datapath string
	if cfg.Env == string(ProdEnv) {
		gin.SetMode(gin.ReleaseMode)
		exe, err := os.Executable()
		if err != nil {
			return "", err
		}
		datapath = filepath.Join(filepath.Dir(exe), cfg.Dir)
	} else if cfg.Env == string(ProdDev) {
		datapath = filepath.Join(".", cfg.Dir)
	} else {
		return "", fmt.Errorf("unexpected env mode")
	}
	return datapath, nil
}
