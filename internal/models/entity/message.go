package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	ChatId     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();foreignkey:ChatId"`
	SenderId   uuid.UUID `gorm:"type:uuid"`
	ReceiverId uuid.UUID `gorm:"type:uuid"`
	Text       string
}
