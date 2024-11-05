package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageEntity struct {
	Id         uuid.UUID `db:"id"`
	ChatId     uuid.UUID `db:"chat_id"`
	EmployeeId uuid.UUID `db:"employee_id"`
	Text       string    `db:"text"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
