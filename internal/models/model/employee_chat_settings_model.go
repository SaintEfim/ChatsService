package model

import "github.com/google/uuid"

type EmployeeChatSettingsModel struct {
	Id          uuid.UUID `json:"id"`
	ChatId      uuid.UUID `json:"chat_id"`
	EmployeeId  uuid.UUID `json:"employee_id"`
	DisplayName string    `json:"display_name"`
}
