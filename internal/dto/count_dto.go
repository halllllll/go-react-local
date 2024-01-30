package dto

// Bind validation bug with int zero value
// https://github.com/gin-gonic/gin/issues/491

type CountInput struct {
	Value *int `json:"count" binding:"required"` // numberは負数に対応しないっぽい
}
