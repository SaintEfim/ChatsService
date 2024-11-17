package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageEntity struct {
	Id          uuid.UUID `json:"id"`
	ChatId      uuid.UUID `json:"chat_id"`
	EmployeeId  uuid.UUID `json:"employee_id"`
	ColleagueId uuid.UUID `json:"colleague_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
