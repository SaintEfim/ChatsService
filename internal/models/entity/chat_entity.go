package entity

import (
	"time"

	"github.com/google/uuid"
)

type ChatEntity struct {
	Id          uuid.UUID   `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	IsGroup     bool        `json:"is_group" db:"is_group"`
	EmployeeIds []uuid.UUID `json:"employee_ids" db:"employee_ids"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}
