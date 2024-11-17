package model

import "github.com/google/uuid"

type EmployeeChatSettingsModel struct {
	Id          uuid.UUID
	ChatId      uuid.UUID
	EmployeeId  uuid.UUID
	DisplayName string
}
