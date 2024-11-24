package interfaces

import (
	"context"
)

type QueryExecutor interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryContext(ctx context.Context, query string, args ...interface{}) (RowsAdapter, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (ResultAdapter, error)
}
