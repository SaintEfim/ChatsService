package model

import "github.com/google/uuid"

type ChatModel struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	IsGroup     bool        `json:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
}
