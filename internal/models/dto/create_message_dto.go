package dto

import "github.com/google/uuid"

type CreateMessageDto struct {
	ChatId      uuid.UUID `json:"chat_id" binding:"required"`
	EmployeeId  uuid.UUID `json:"employee_id" binding:"required"`
	ColleagueId uuid.UUID `json:"colleague_id" binding:"required"`
	Text        string    `json:"text" binding:"required"`
}
