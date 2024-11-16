package dto

import "github.com/google/uuid"

type ChatDetailDTO struct {
	Id          uuid.UUID   `json:"id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
}
