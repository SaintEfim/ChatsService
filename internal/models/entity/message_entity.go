package entity

import "github.com/google/uuid"

type MessageEntity struct {
	Id         uuid.UUID `db:"id"`
	ChatId     uuid.UUID `db:"chatId"`
	EmployeeId uuid.UUID `db:"employeeId"`
	Text       string    `db:"text"`
}
