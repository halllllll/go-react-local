package router

import (
	"database/sql"
	"log/slog"
	"net/http"
	"sample/go-react-local-app/internal/common/di"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

func SetCountRoutes(r *gin.Engine, db *sql.DB, logger *slog.Logger) {
	countCtrl := di.InitCount(db, logger)
	api := r.Group("/api")
	{
		api.POST("/count", countCtrl.AddCount)
		api.GET("/count/:id", countCtrl.GetCount)
	}
	r.GET("/health", healthHandler)

}
