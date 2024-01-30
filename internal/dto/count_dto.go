package dto

type CountInput struct {
	Value int `json:"value" binding:"required"`
}
