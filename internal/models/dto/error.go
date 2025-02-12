package dto

type Error struct {
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	StackTrace  string `json:"stackTrace,omitempty"`
}
