package executor

import (
	"ChatsService/config"
	"context"
	"database/sql"

	"ChatsService/internal/models/interfaces"

	"github.com/jmoiron/sqlx"
)

func PostgresConnect(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DataBase.DriverName, cfg.DataBase.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

type SQLXExecutor struct {
	db *sqlx.DB
}

func NewSQLXExecutor(db *sqlx.DB) interfaces.QueryExecutor {
	return &SQLXExecutor{db: db}
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
