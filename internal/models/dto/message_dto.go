package dto

import (
	"time"

	"github.com/google/uuid"
)

type MessageDto struct {
	Id         uuid.UUID `json:"id"`
	ChatId     uuid.UUID `json:"chatId"`
	EmployeeId uuid.UUID `json:"employeeId"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
