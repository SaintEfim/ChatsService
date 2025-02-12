package dto

type MessageUpdate struct {
	Text string `json:"text" binding:"required"`
}
