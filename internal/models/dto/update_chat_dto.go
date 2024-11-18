package dto

import "github.com/google/uuid"

type UpdateChatDto struct {
	Name        string      `json:"name"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
}
