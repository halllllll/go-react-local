package controller

import (
	"net/http"
	"sample/go-react-local-app/internal/dto"
	"sample/go-react-local-app/internal/service"

	"github.com/gin-gonic/gin"
)

type CountControler interface {
	AddCount(ctx *gin.Context)
}

type counter struct {
	service service.CountServicer
}

func NewCountController(service service.CountServicer) CountControler {
	return &counter{
		service: service,
	}
}

func (c *counter) AddCount(ctx *gin.Context) {
	var input dto.CountInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if err := c.service.Set(ctx, input.Value); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"newCount": input.Value,
	})
}