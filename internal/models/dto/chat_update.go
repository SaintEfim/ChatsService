package dto

import "github.com/google/uuid"

type ChatUpdate struct {
	Name           string      `json:"name"`
	ParticipantIds []uuid.UUID `json:"participant_ids" binding:"required"`
}
