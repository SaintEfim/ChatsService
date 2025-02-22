package dto

import "github.com/google/uuid"

type Employee struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
}
