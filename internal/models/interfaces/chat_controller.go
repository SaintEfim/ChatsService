package interfaces

import (
	"context"

	"github.com/google/uuid"

	"ChatsService/internal/models/dto"
)

type ChatController interface {
	Get(ctx context.Context) ([]*dto.Chat, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error)
	Create(ctx context.Context, model *dto.ChatDetail) (*dto.ChatDetail, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, model *dto.ChatDetail) error
}
