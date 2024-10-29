package psql

const (
	retrieveAllMessagesByChatId = `SELECT id, employee_id, text FROM messages WHERE chat_id = $1`
	retrieveMessageById         = `SELECT id, chat_id, employee_id, text FROM messages WHERE id = $1`
	createMessage               = `INSERT INTO messages (id, chat_id, employee_id, text) VALUES ($1, $2, $3, $4)`
	deleteMessage               = `DELETE FROM messages WHERE id = $1`
	updateMessage               = `UPDATE messages SET text = $1 WHERE id = $2`
)
