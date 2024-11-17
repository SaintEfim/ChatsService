package entity

import (
	"time"

	"github.com/google/uuid"
)

type ChatEntity struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	IsGroup     bool        `json:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
