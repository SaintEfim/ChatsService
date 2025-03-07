package dto

import "github.com/google/uuid"

type Chat struct {
	Id           uuid.UUID     `json:"id" binding:"required"`
	Name         string        `json:"name"`
	Participants []Participant `json:"participants" binding:"required"`
}
