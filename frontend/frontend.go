package frontend

import (
	"embed"
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

var (
	//go:embed dist
	dist embed.FS

	devPort int = 5173 // vite dev server port
)

func RegisterHandlers(router *gin.Engine) {
	if os.Getenv("ENV") == "dev" {
		log.Println("Running on dev mode")
		setupProxy(router)
	} else {
		gin.SetMode(gin.ReleaseMode)

		// 静的ファイル(embed対応)へのアクセスを設定
		router.Use(static.ServeEmbed("dist", dist))

	}
}

func setupProxy(r *gin.Engine) {
	targetURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", devPort))
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			proxy.ServeHTTP(c.Writer, c.Request)
		} else {
			c.Next()
		}
	})

}
