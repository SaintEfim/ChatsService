package dto

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID `json:"id" binding:"required"`
	ChatId     uuid.UUID `json:"chat_id" binding:"required"`
	SenderId   uuid.UUID `json:"sender_id"  binding:"required"`
	ReceiverId uuid.UUID `json:"receiver_id"  binding:"required"`
	Text       string    `json:"text"  binding:"required"`
	CreatedAt  time.Time `json:"created_at"  binding:"required"`
}
