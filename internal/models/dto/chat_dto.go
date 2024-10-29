package dto

import "github.com/google/uuid"

type ChatDTO struct {
	Id           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	EmployeesIds []uuid.UUID `json:"employeesIds"`
}
