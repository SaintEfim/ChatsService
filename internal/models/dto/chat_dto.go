package dto

import "github.com/google/uuid"

type ChatDTO struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	EmployeeIds []uuid.UUID `json:"employeesIds"`
}
