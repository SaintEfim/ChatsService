package dto

import "github.com/google/uuid"

type ChatDetail struct {
	Id        uuid.UUID  `json:"id" binding:"required"`
	Name      string     `json:"name"`
	IsGroup   bool       `json:"is_group"`
	Employees []Employee `json:"employees"`
}
