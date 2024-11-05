package interfaces

import "context"

type Controller[T any] interface {
	Get(ctx context.Context) ([]*T, error)
	GetOneById(ctx context.Context, id string) (*T, error)
	Create(ctx context.Context, user *T) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, user *T) error
}
