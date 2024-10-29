package interfaces

import (
	"context"

	"ChatsService/internal/models/entity"

	"github.com/google/uuid"
)

type MessageRepository interface {
	GetAllByChatId(ctx context.Context, chatId uuid.UUID) ([]*entity.MessageEntity, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.MessageEntity, error)
	Create(ctx context.Context, message *entity.MessageEntity) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, message *entity.MessageEntity) error
}
