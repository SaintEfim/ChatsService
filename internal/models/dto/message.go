package dto

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID `json:"id" binding:"required"`
	ChatId     uuid.UUID `json:"chatId" binding:"required"`
	SenderId   uuid.UUID `json:"senderId"  binding:"required"`
	ReceiverId uuid.UUID `json:"receiverId"  binding:"required"`
	Text       string    `json:"text"  binding:"required"`
	CreatedAt  time.Time `json:"createdAt"  binding:"required"`
}
