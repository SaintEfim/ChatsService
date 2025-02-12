package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type Repository[TEntity any] interface {
	Get(ctx context.Context) ([]*TEntity, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*TEntity, error)
	Create(ctx context.Context, entity *TEntity) (*TEntity, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, entity *TEntity) error
}
