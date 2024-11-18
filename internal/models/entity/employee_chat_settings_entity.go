package entity

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeChatSettingsEntity struct {
	Id          uuid.UUID `db:"id"`
	ChatId      uuid.UUID `db:"chat_id"`
	EmployeeId  uuid.UUID `db:"employee_id"`
	DisplayName string    `db:"display_name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
