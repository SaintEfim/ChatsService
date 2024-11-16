package interfaces

import (
	"context"
	"github.com/google/uuid"
)

type Controller[T any] interface {
	Get(ctx context.Context) ([]*T, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*T, error)
	Create(ctx context.Context, user *T) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, user *T) error
}
