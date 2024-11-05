package dto

import "github.com/google/uuid"

type CreateChatDto struct {
	Name        string      `json:"name"`
	EmployeeIds []uuid.UUID `json:"employeesIds"`
}
