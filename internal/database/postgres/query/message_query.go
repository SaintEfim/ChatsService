package query

import (
	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
)

const (
	getAllMessages = `SELECT id, chat_id, employee_id, colleague_id, text, created_at FROM messages`
	getMessageById = `SELECT id, chat_id, employee_id, colleague_id, text FROM messages WHERE id = $1`
	createMessage  = `INSERT INTO messages (id, chat_id, employeeId, colleague_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteMessage  = `DELETE FROM messages WHERE id = $1`
	updateMessage  = `UPDATE messages SET text = $1, updated_at = NOW() WHERE id = $2`
)

type MessageQuery struct{}

func NewMessageQuery() interfaces.Query[dto.Message] {
	return &MessageQuery{}
}

func (c *MessageQuery) Get() string {
	return getAllMessages
}

func (c *MessageQuery) GetOneById() string {
	return getMessageById
}

func (c *MessageQuery) Create() string {
	return createMessage
}

func (c *MessageQuery) Delete() string {
	return deleteMessage
}

func (c *MessageQuery) Update() string {
	return updateMessage
}
