package repository

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type EmployeeChatSettingsRepository struct {
	db    interfaces.QueryExecutor
	query interfaces.Query[entity.EmployeeChatSettingsEntity]
}

func NewEmployeeChatSettingsRepository(db interfaces.QueryExecutor, query interfaces.Query[entity.EmployeeChatSettingsEntity]) interfaces.Repository[entity.EmployeeChatSettingsEntity] {
	return &EmployeeChatSettingsRepository{
		db:    db,
		query: query,
	}
}

func (r *EmployeeChatSettingsRepository) Get(ctx context.Context) ([]*entity.EmployeeChatSettingsEntity, error) {
	employeeChatSettings := make([]*entity.EmployeeChatSettingsEntity, 0)

	err := r.db.SelectContext(ctx, &employeeChatSettings, r.query.Get())
	if err != nil {
		return nil, err
	}

	return employeeChatSettings, nil
}

func (r *EmployeeChatSettingsRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.EmployeeChatSettingsEntity, error) {
	employeeChatSettings := &entity.EmployeeChatSettingsEntity{}

	if err := r.db.GetContext(ctx, &employeeChatSettings, r.query.Get(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return employeeChatSettings, nil
}

func (r *EmployeeChatSettingsRepository) Create(ctx context.Context, employeeChatSettings *entity.EmployeeChatSettingsEntity) (uuid.UUID, error) {
	employeeChatSettings.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, r.query.Create(),
		employeeChatSettings.Id,
		employeeChatSettings.ChatId,
		employeeChatSettings.EmployeeId,
		employeeChatSettings.DisplayName)
	if err != nil {
		return uuid.Nil, err
	}

	return employeeChatSettings.Id, nil
}

func (r *EmployeeChatSettingsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, r.query.Delete(), id)
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
	result, err := r.db.ExecContext(ctx, r.query.Update(),
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
