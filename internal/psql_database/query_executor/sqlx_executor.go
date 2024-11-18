package query_executor

import (
	"context"

	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SQLXExecutor struct {
	db *sqlx.DB
}

func (e *SQLXExecutor) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return e.db.SelectContext(ctx, dest, query, args...)
}

func (e *SQLXExecutor) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return e.db.GetContext(ctx, dest, query, args...)
}

func (e *SQLXExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return e.db.QueryContext(ctx, query, args...)
}

func (e *SQLXExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return e.db.ExecContext(ctx, query, args...)
}
