package psql

const (
	retrieveAllChats = `SELECT id, name, employees_ids FROM chats`
	retrieveChatById = `SELECT id, name, employees_ids FROM chats WHERE id = $1`
	createChat       = `INSERT INTO chats (id, name, employees_ids) VALUES ($1, $2, $3)`
	deleteChat       = `DELETE FROM chats WHERE id = $1`
	updateChat       = `UPDATE chats SET name = $1, employees_ids = $2 WHERE id = $3`
)
