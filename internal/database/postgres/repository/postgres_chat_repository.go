package repository

import (
	"context"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/repository"

	"github.com/google/uuid"
)

type PostgresChatRepository struct {
	repo  interfaces.Repository[entity.ChatEntity]
	query interfaces.Query[entity.ChatEntity]
}

func NewPostgresChatRepository(exec interfaces.QueryExecutor, query interfaces.Query[entity.ChatEntity]) interfaces.Repository[entity.ChatEntity] {
	baseRepo := repository.NewChatRepository(exec, query)
	return &PostgresChatRepository{repo: baseRepo}
}

func (r *PostgresChatRepository) Get(ctx context.Context) ([]*entity.ChatEntity, error) {
	return r.repo.Get(ctx)
}

func (r *PostgresChatRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.ChatEntity, error) {
	return r.repo.GetOneById(ctx, id)
}

func (r *PostgresChatRepository) Create(ctx context.Context, chat *entity.ChatEntity) (uuid.UUID, error) {
	return r.repo.Create(ctx, chat)
}

func (r *PostgresChatRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.repo.Delete(ctx, id)
}

func (r *PostgresChatRepository) Update(ctx context.Context, id uuid.UUID, chat *entity.ChatEntity) error {
	return r.repo.Update(ctx, id, chat)
}
