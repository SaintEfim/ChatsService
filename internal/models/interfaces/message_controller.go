package interfaces

import (
	"context"

	"github.com/google/uuid"

	"ChatsService/internal/models/dto"
)

type MessageController interface {
	Get(ctx context.Context) ([]*dto.Message, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error)
	Create(ctx context.Context, model *dto.Message) (*dto.Message, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, model *dto.Message) error
}
