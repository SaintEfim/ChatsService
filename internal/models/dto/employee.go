package dto

import "github.com/google/uuid"

type Participant struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Patronymic string    `json:"patronymic"`
}
