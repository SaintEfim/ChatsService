package entity

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeChatSettingsEntity struct {
	Id          uuid.UUID `json:"id"`
	ChatId      uuid.UUID `json:"chat_id"`
	EmployeeId  uuid.UUID `json:"employee_id"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
