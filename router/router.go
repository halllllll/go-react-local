package router

import (
	"context"
	"database/sql"
	"net/http"
	"sample/go-react-local-app/handler"

	"github.com/gin-gonic/gin"
)



func healthHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func SetRoutes(r *gin.Engine, db *sql.DB, ctx context.Context){
	api := r.Group("/api")
	apiRegister := &handler.ApiRegister{
		DB: db,
		Ctx: ctx,
	}

	{
		api.GET("/hello", apiRegister.Hello)
		api.POST("/hello", apiRegister.HelloPost)
	}
	r.GET("/health", healthHandler)
}