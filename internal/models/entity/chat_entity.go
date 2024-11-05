package entity

import (
	"time"

	"github.com/google/uuid"
)

type ChatEntity struct {
	Id          uuid.UUID   `db:"id"`
	Name        string      `db:"name"`
	EmployeeIds []uuid.UUID `db:"employee_ids"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}
