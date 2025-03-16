package interfaces

import (
	"context"

	"ChatsService/internal/models/dto"

	"github.com/google/uuid"
)

type ChatController interface {
	Get(ctx context.Context) ([]*dto.Chat, error)
	GetChatsByUserId(ctx context.Context, userId uuid.UUID) ([]*dto.Chat, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error)
	Create(ctx context.Context, model *dto.ChatCreate) (*dto.CreateAction, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, model *dto.ChatUpdate) error
}
