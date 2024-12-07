package dto

import "github.com/google/uuid"

type ChatCreate struct {
	Name        string      `json:"name" binding:"required"`
	IsGroup     bool        `json:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
}
