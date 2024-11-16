package dto

type ErrorDto struct {
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
}
