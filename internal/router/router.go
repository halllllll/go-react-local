package router

import (
	"net/http"
	"sample/go-react-local-app/internal/controller"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// func SetRoutes(r *gin.Engine, db *sql.DB, ctx context.Context) {
func SetRoutes(r *gin.Engine, ctrl controller.CountControler) {
	api := r.Group("/api")

	// apiRegister := &handler.ApiRegister{
	// 	Repo: repository.Repository{DB: db, Ctx: ctx},
	// }
	{
		// api.GET("/count", apiRegister.GetCount)
		// api.POST("/count", apiRegister.SetCount)
		api.GET("/count", ctrl.AddCount)
	}
	r.GET("/health", healthHandler)
}
