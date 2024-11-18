package query

const (
	GetAllEmployeeChatSettings  = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings`
	GetEmployeeChatSettingsById = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings WHERE id = $1`
	CreateEmployeeChatSettings  = `INSERT INTO employee_chat_settings (id, chat_id, employee_id, display_name, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	DeleteEmployeeChatSettings  = `DELETE FROM employee_chat_settings WHERE id = $1`
	UpdateEmployeeChatSettings  = `UPDATE employee_chat_settings SET display_name = $1, updated_at = NOW() WHERE id = $2`
)
