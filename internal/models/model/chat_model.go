package model

import "github.com/google/uuid"

type ChatModel struct {
	Id          uuid.UUID
	Name        string
	IsGroup     bool
	EmployeeIds []uuid.UUID
}
