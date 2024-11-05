package model

import "github.com/google/uuid"

type ChatModel struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	EmployeeIds []uuid.UUID `json:"employeesIds"`
}
