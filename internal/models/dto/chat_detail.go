package dto

import "github.com/google/uuid"

type ChatDetail struct {
	Id          uuid.UUID   `json:"id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	IsGroup     bool        `json:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
}
