package dto

import "github.com/google/uuid"

type ChatDto struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}
