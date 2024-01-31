package controller

import (
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
	GetCounts(ctx *gin.Context)
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

	count, err := c.service.Get(ctx, int(countId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":      countId,
		"value":   count.Val,
		"created": count.Created,
		"updated": count.Updated,
	})
	return
}

// GetCounts implements CountControler.
func (c *counter) GetCounts(ctx *gin.Context) {
	counts, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": counts,
	})
}
