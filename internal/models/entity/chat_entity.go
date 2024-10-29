package entity

import "github.com/google/uuid"

type ChatEntity struct {
	Id           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	EmployeesIds []uuid.UUID `json:"employeesIds"`
}
