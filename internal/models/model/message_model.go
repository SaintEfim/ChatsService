package model

import (
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id          uuid.UUID `json:"id"`
	ChatId      uuid.UUID `json:"chatId"`
	EmployeeId  uuid.UUID `json:"employeeId"`
	ColleagueId uuid.UUID `json:"colleague_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}
