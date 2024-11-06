package dto

import "github.com/google/uuid"

type CreateChatDto struct {
	Name        string      `json:"name"`
	IsGroup     bool        `db:"is_group"`
	EmployeeIds []uuid.UUID `json:"employeesIds"`
}
