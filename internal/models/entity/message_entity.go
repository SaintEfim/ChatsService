package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageEntity struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ChatId      uuid.UUID `json:"chat_id" db:"chat_id"`
	EmployeeId  uuid.UUID `json:"employee_id" db:"employee_id"`
	ColleagueId uuid.UUID `json:"colleague_id" db:"colleague_id"`
	Text        string    `json:"text" db:"text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
