package entity

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeChatSettingsEntity struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ChatId      uuid.UUID `json:"chat_id" db:"chat_id"`
	EmployeeId  uuid.UUID `json:"employee_id" db:"employee_id"`
	DisplayName string    `json:"display_name" db:"display_name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
