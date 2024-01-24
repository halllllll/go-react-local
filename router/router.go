package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "yoyo~!"})
}

func healthHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func SetRoutes(r *gin.Engine){
	api := r.Group("/api")

	{
		api.GET("/hello", apiHello)
		
	}
	r.GET("/health", healthHandler)
}