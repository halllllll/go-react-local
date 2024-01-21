package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sample/go-react-local-app/frontend"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/browser"
	sloggin "github.com/samber/slog-gin"
)

var (
	port int = 3056
)

func ApiHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "yoyo~!"})
}

func main(){
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	router := gin.Default()
	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

	
	api := router.Group("/api")
	
	{
		api.GET("/hello", ApiHello)
	}
	
	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d", port)); err != nil{
		log.Println(err)
	}
	frontend.RegisterHandlers(router)

	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}