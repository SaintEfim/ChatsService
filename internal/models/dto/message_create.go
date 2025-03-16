package dto

import (
	"time"

	"github.com/google/uuid"
)

type MessageCreate struct {
	ChatId     uuid.UUID `json:"chatId" binding:"required"`
	SenderId   uuid.UUID `json:"senderId" binding:"required"`
	ReceiverId uuid.UUID `json:"receiverId" binding:"required"`
	Text       string    `json:"text" binding:"required"`
	CreateAt   time.Time `json:"createAt" binding:"required"`
}
