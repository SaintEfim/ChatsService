package dto

import "github.com/google/uuid"

type ChatCreate struct {
	ParticipantIds []uuid.UUID `json:"participantIds" binding:"required"`
}
