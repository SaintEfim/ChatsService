package dto

import "github.com/google/uuid"

type ChatCreate struct {
	Name           string      `json:"name"`
	IsGroup        bool        `json:"is_group"`
	ParticipantIds []uuid.UUID `json:"participant_ids" binding:"required"`
}
