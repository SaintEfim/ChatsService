package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Id          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string
	IsGroup     bool        `gorm:"default:false"`
	EmployeeIds []uuid.UUID `gorm:"type:uuid[]"`
	Messages    []Message
}
