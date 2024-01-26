package router

import (
	"context"
	"database/sql"
	"net/http"
	"sample/go-react-local-app/handler"
	"sample/go-react-local-app/repository"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func SetRoutes(r *gin.Engine, db *sql.DB, ctx context.Context) {
	api := r.Group("/api")

	apiRegister := &handler.ApiRegister{
		Repo: repository.Repository{DB: db, Ctx: ctx},
	}
	{
		api.GET("/count", apiRegister.GetCount)
		api.POST("/count", apiRegister.SetCount)
	}
	r.GET("/health", healthHandler)
}
