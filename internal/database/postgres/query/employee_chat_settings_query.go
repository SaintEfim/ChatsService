package query

import (
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
)

const (
	getAllEmployeeChatSettings  = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings`
	getEmployeeChatSettingsById = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings WHERE id = $1`
	createEmployeeChatSettings  = `INSERT INTO employee_chat_settings (id, chat_id, employee_id, display_name, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteEmployeeChatSettings  = `DELETE FROM employee_chat_settings WHERE id = $1`
	updateEmployeeChatSettings  = `UPDATE employee_chat_settings SET display_name = $1, updated_at = NOW() WHERE id = $2`
)

type EmployeeChatSettingsQuery struct{}

func NewEmployeeChatSettingsQueryQuery() interfaces.Query[entity.EmployeeChatSettingsEntity] {
	return &EmployeeChatSettingsQuery{}
}

func (c *EmployeeChatSettingsQuery) Get() string {
	return getAllEmployeeChatSettings
}

func (c *EmployeeChatSettingsQuery) GetOneById() string {
	return getEmployeeChatSettingsById
}

func (c *EmployeeChatSettingsQuery) Create() string {
	return createEmployeeChatSettings
}

func (c *EmployeeChatSettingsQuery) Delete() string {
	return deleteEmployeeChatSettings
}

func (c *EmployeeChatSettingsQuery) Update() string {
	return updateEmployeeChatSettings
}
