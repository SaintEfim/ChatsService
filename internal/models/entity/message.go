package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func (Message) TableName() string {
	return "Messages"
}

type Message struct {
	Id         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:Id"`
	CreatedAt  time.Time      ``
	UpdatedAt  time.Time      ``
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ChatId     uuid.UUID      `gorm:"type:uuid;column:ChatId"`
	SenderId   uuid.UUID      `gorm:"type:uuid;column:SenderId"`
	ReceiverId uuid.UUID      `gorm:"type:uuid;column:ReceiverId"`
	Text       string         `gorm:"column:Text"`
}
