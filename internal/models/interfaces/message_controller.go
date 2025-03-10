package interfaces

import (
	"context"

	"ChatsService/internal/models/dto"

	"github.com/google/uuid"
)

type MessageController interface {
	Get(ctx context.Context) ([]*dto.Message, error)
	GetMessagesByChatId(ctx context.Context, chatId uuid.UUID) ([]*dto.Message, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error)
	Create(ctx context.Context, model *dto.MessageCreate) (*dto.Message, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, model *dto.MessageUpdate) error
}
