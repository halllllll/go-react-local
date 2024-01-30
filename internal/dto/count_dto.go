package dto

type CountInput struct {
	Value int `json:"count" binding:"required"`
}
