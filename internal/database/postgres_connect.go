package database

import (
	"ChatsService/config"
	"context"

	"github.com/jmoiron/sqlx"
)

func PostgresConnect(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionStrings.ServiceDb)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
