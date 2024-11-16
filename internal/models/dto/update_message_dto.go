package dto

type UpdateMessageDto struct {
	Text string `json:"text" binding:"required"`
}
