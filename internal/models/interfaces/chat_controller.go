package interfaces

import (
	"ChatsService/internal/models/dto"
	"context"

	"github.com/google/uuid"
)

type ChatController interface {
	Get(ctx context.Context) ([]*dto.Chat, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error)
	Create(ctx context.Context, model *dto.ChatDetail) (*dto.ChatDetail, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, model *dto.ChatDetail) error
}
