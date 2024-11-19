package query

import (
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
)

const (
	getAllChats = `SELECT id, name, is_group, employee_ids FROM chats`
	getChatById = `SELECT id, name, is_group, employee_ids FROM chats WHERE id = $1`
	createChat  = `INSERT INTO chats (id, name, is_group, employee_ids, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteChat  = `DELETE FROM chats WHERE id = $1`
	updateChat  = `UPDATE chats SET name = $1, employee_ids = $2, updated_at = NOW() WHERE id = $3`
)

type ChatQuery struct{}

func NewChatQuery() interfaces.Query[entity.ChatEntity] {
	return &ChatQuery{}
}

func (c *ChatQuery) Get() string {
	return getAllChats
}

func (c *ChatQuery) GetOneById() string {
	return getChatById
}

func (c *ChatQuery) Create() string {
	return createChat
}

func (c *ChatQuery) Delete() string {
	return deleteChat
}

func (c *ChatQuery) Update() string {
	return updateChat
}
