package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type Repository[TEntity, TDetail, TCreate, TUpdate any] interface {
	Get(ctx context.Context) ([]*TEntity, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*TDetail, error)
	Create(ctx context.Context, entity *TCreate) (*TDetail, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, entity *TUpdate) error
}
