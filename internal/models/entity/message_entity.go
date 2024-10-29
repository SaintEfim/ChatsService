package entity

import "github.com/google/uuid"

type MessageEntity struct {
	Id         uuid.UUID `json:"id"`
	ChatId     uuid.UUID `json:"chatId"`
	EmployeeId uuid.UUID `json:"employeeId"`
	Text       string    `json:"text"`
}
