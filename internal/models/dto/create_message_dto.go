package dto

import "github.com/google/uuid"

type CreateMessageDto struct {
	ChatId     uuid.UUID `json:"chatId"`
	EmployeeId uuid.UUID `json:"employeeId"`
	Text       string    `json:"text"`
}
