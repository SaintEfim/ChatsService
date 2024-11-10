package repository

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	retrieveAllEmployeeChatSettings  = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings`
	retrieveEmployeeChatSettingsById = `SELECT id, chat_id, employee_id, display_name FROM employee_chat_settings WHERE id = $1`
	createEmployeeChatSettings       = `INSERT INTO employee_chat_settings (id, chat_id, employee_id, display_name, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteEmployeeChatSettings       = `DELETE FROM employee_chat_settings WHERE id = $1`
	updateEmployeeChatSettings       = `UPDATE employee_chat_settings SET display_name = $1, updated_at = NOW() WHERE id = $2`
)

type EmployeeChatSettingsRepository struct {
	db *sqlx.DB
}

func NewEmployeeChatSettingsRepository(db *sqlx.DB) interfaces.Repository[entity.EmployeeChatSettingsEntity] {
	return &EmployeeChatSettingsRepository{
		db: db,
	}
}

func (r *EmployeeChatSettingsRepository) Get(ctx context.Context) ([]*entity.EmployeeChatSettingsEntity, error) {
	employeeChatSettings := make([]*entity.EmployeeChatSettingsEntity, 0)

	err := r.db.SelectContext(ctx, &employeeChatSettings, retrieveAllEmployeeChatSettings)
	if err != nil {
		return nil, err
	}

	return employeeChatSettings, nil
}

func (r *EmployeeChatSettingsRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.EmployeeChatSettingsEntity, error) {
	employeeChatSettings := &entity.EmployeeChatSettingsEntity{}

	if err := r.db.GetContext(ctx, &employeeChatSettings, retrieveEmployeeChatSettingsById, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return employeeChatSettings, nil
}

func (r *EmployeeChatSettingsRepository) Create(ctx context.Context, employeeChatSettings *entity.EmployeeChatSettingsEntity) error {
	employeeChatSettings.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createEmployeeChatSettings,
		employeeChatSettings.Id,
		employeeChatSettings.ChatId,
		employeeChatSettings.EmployeeId,
		employeeChatSettings.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeChatSettingsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteEmployeeChatSettings, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (r *EmployeeChatSettingsRepository) Update(ctx context.Context, id uuid.UUID, employeeChatSettings *entity.EmployeeChatSettingsEntity) error {
	result, err := r.db.ExecContext(ctx, updateEmployeeChatSettings,
		employeeChatSettings.DisplayName,
		id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
