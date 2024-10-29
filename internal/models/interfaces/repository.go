package interfaces

import (
	"ChatsService/internal/models/entity"
	"context"
	"github.com/google/uuid"
)

type ChatRepository interface {
	Get(ctx context.Context) ([]*entity.ChatEntity, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.ChatEntity, error)
	Create(ctx context.Context, entity *entity.ChatEntity) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, entity *entity.ChatEntity) error
}
