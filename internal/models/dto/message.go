package dto

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id          uuid.UUID `json:"id" db:"id" binding:"required"`
	ChatId      uuid.UUID `json:"chat_id" db:"chat_id" binding:"required"`
	EmployeeId  uuid.UUID `json:"employee_id" db:"employee_id" binding:"required"`
	ColleagueId uuid.UUID `json:"colleague_id" db:"colleague_id" binding:"required"`
	Text        string    `json:"text" db:"text" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" binding:"required"`
}
