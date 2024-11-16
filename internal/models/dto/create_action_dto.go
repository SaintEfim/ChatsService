package dto

import "github.com/google/uuid"

type CreateActionDto struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
