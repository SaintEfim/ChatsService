package query

const (
	GetAllChats = `SELECT id, name, is_group, employee_ids FROM chats`
	GetChatById = `SELECT id, name, is_group, employee_ids FROM chats WHERE id = $1`
	CreateChat  = `INSERT INTO chats (id, name, is_group, employee_ids, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	DeleteChat  = `DELETE FROM chats WHERE id = $1`
	UpdateChat  = `UPDATE chats SET name = $1, employee_ids = $2, updated_at = NOW() WHERE id = $3`
)
