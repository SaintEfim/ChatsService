package dto

import (
	"github.com/google/uuid"
	"time"
)

type MessageCreate struct {
	ChatId     uuid.UUID `json:"chat_id" binding:"required"`
	SenderId   uuid.UUID `json:"sender_id" binding:"required"`
	ReceiverId uuid.UUID `json:"receiver_id" binding:"required"`
	Text       string    `json:"text" binding:"required"`
	CreateAt   time.Time `json:"create_at" binding:"required"`
}
