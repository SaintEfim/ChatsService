package repository

import (
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/repository"
	"context"
	"github.com/google/uuid"
)

type PostgresMessageRepository struct {
	repo  interfaces.Repository[entity.MessageEntity]
	query interfaces.Query[entity.MessageEntity]
}

func NewPostgresMessageRepository(exec interfaces.QueryExecutor, query interfaces.Query[entity.MessageEntity]) interfaces.Repository[entity.MessageEntity] {
	baseRepo := repository.NewMessageRepository(exec, query)
	return &PostgresMessageRepository{repo: baseRepo}
}

func (r *PostgresMessageRepository) Get(ctx context.Context) ([]*entity.MessageEntity, error) {
	return r.repo.Get(ctx)
}

func (r *PostgresMessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.MessageEntity, error) {
	return r.repo.GetOneById(ctx, id)
}

func (r *PostgresMessageRepository) Create(ctx context.Context, chat *entity.MessageEntity) (uuid.UUID, error) {
	return r.repo.Create(ctx, chat)
}

func (r *PostgresMessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.repo.Delete(ctx, id)
}

func (r *PostgresMessageRepository) Update(ctx context.Context, id uuid.UUID, chat *entity.MessageEntity) error {
	return r.repo.Update(ctx, id, chat)
}
