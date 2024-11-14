package psql

import (
	"context"

	"ChatsService/config"

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
