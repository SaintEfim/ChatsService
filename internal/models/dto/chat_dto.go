package dto

import "github.com/google/uuid"

type ChatDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
