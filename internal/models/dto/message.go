package dto

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	ChatId      uuid.UUID `json:"chat_id" binding:"required"`
	EmployeeId  uuid.UUID `json:"employee_id"  binding:"required"`
	ColleagueId uuid.UUID `json:"colleague_id"  binding:"required"`
	Text        string    `json:"text"  binding:"required"`
	CreatedAt   time.Time `json:"created_at"  binding:"required"`
}
