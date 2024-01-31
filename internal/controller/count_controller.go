package controller

import (
	"fmt"
	"log/slog"
	"net/http"
	"sample/go-react-local-app/internal/common/dto"
	"sample/go-react-local-app/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CountControler interface {
	AddCount(ctx *gin.Context)
	GetCount(ctx *gin.Context)
}

type counter struct {
	service service.CountServicer
	logger  *slog.Logger
}

func NewCountController(service service.CountServicer, logger *slog.Logger) CountControler {
	return &counter{
		service: service,
		logger:  logger,
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
	if err := c.service.Set(ctx, *input.Value); err != nil {
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
	return
}

func (c *counter) GetCount(ctx *gin.Context) {
	countId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	c.logger.Info(fmt.Sprintf("id(dummy): %d", countId))

	ctx.JSON(http.StatusOK, gin.H{
		"id":    countId,
		"value": 10000000,
	})
	return
}
