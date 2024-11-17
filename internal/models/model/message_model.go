package model

import (
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id          uuid.UUID
	ChatId      uuid.UUID
	EmployeeId  uuid.UUID
	ColleagueId uuid.UUID
	Text        string
	CreatedAt   time.Time
}
