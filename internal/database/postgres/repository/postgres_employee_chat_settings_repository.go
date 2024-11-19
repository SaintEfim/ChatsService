package repository

import (
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/repository"
	"context"
	"github.com/google/uuid"
)

type PostgresEmployeeChatSettingsRepository struct {
	repo  interfaces.Repository[entity.EmployeeChatSettingsEntity]
	query interfaces.Query[entity.EmployeeChatSettingsEntity]
}

func NewPostgresEmployeeChatSettingsRepository(exec interfaces.QueryExecutor, query interfaces.Query[entity.EmployeeChatSettingsEntity]) interfaces.Repository[entity.EmployeeChatSettingsEntity] {
	baseRepo := repository.NewEmployeeChatSettingsRepository(exec, query)
	return &PostgresEmployeeChatSettingsRepository{repo: baseRepo}
}

func (r *PostgresEmployeeChatSettingsRepository) Get(ctx context.Context) ([]*entity.EmployeeChatSettingsEntity, error) {
	return r.repo.Get(ctx)
}

func (r *PostgresEmployeeChatSettingsRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.EmployeeChatSettingsEntity, error) {
	return r.repo.GetOneById(ctx, id)
}

func (r *PostgresEmployeeChatSettingsRepository) Create(ctx context.Context, chat *entity.EmployeeChatSettingsEntity) (uuid.UUID, error) {
	return r.repo.Create(ctx, chat)
}

func (r *PostgresEmployeeChatSettingsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.repo.Delete(ctx, id)
}

func (r *PostgresEmployeeChatSettingsRepository) Update(ctx context.Context, id uuid.UUID, chat *entity.EmployeeChatSettingsEntity) error {
	return r.repo.Update(ctx, id, chat)
}
