package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"sample/go-react-local-app/frontend"
	"sample/go-react-local-app/internal/common/config"
	"sample/go-react-local-app/internal/db"
	"sample/go-react-local-app/internal/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	sloggin "github.com/samber/slog-gin"
)

// info: Makefileのbuild参照。`make build`時にプロダクション用の値をセットする
var AppMode string

func main() {
	// go run(make dev)時と実行ファイルの実行時でカレントディレクトリを切り替える
	if err := os.Setenv("ENV", AppMode); err != nil {
		log.Fatal(err)
	}
	if os.Getenv("ENV") == string(config.EnvProd) {
		os.Setenv("GIN_MODE", "release")
		gin.SetMode(gin.ReleaseMode)
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
	datapath, err := cfg.CheckEnv()
	if err != nil {
		return err
	}

	db, cleanupdb, err := db.NewDB(ctx, datapath)
	defer cleanupdb()
	if err != nil {
		return err
	}
	logger, cleanuplog, err := cfg.CreateAppLog(datapath)
	if err != nil {
		return err
	}
	defer cleanuplog()
	r := gin.Default()
	r.ContextWithFallback = true

	// middlewares
	r.Use(sloggin.New(logger[string(config.GinLog)]))
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

	router.SetCountRoutes(r, db, logger[string(config.AppLog)])
	if cfg.Env == config.EnvProd {
		// フロントの埋め込みファイル参照はルーティング設定のあとにしないと404が返る
		frontend.RegisterHandlers(r)
		if err := browser.OpenURL(fmt.Sprintf("http://%s", cfg.Address)); err != nil {
			logger[string(config.AppLog)].Error(err.Error())
		}
	} else if cfg.Env == config.EnvDev {
		frontend.SetupProxy(r)
	}

	err = r.Run(fmt.Sprintf(":%d", cfg.Port))
	return err
}
