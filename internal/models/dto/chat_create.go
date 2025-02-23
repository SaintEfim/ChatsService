package dto

import "github.com/google/uuid"

type ChatCreate struct {
	Name        string      `json:"name"`
	IsGroup     bool        `json:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids" binding:"required"`
}
