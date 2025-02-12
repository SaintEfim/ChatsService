package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type Controller[TModel, TDetail, TCreate, TUpdate any] interface {
	Get(ctx context.Context) ([]*TModel, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*TDetail, error)
	Create(ctx context.Context, model *TCreate) (*TDetail, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, user *TUpdate) error
}
