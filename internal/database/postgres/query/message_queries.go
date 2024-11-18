package query

const (
	GetAllMessages = `SELECT id, chat_id, employee_id, colleague_id, text FROM messages`
	GetMessageById = `SELECT id, chat_id, employee_id, colleague_id, text FROM messages WHERE id = $1`
	CreateMessage  = `INSERT INTO messages (id, chat_id, employeeId, colleague_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	DeleteMessage  = `DELETE FROM messages WHERE id = $1`
	UpdateMessage  = `UPDATE messages SET text = $1, updated_at = NOW() WHERE id = $2`
)
