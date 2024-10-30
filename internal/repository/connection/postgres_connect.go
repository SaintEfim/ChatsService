package connection

import (
	"ChatsService/config"

	"database/sql"
)

func PostgresConnect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionStrings.ServiceDb)

	if err != nil {
		return nil, err
	}

	return db, nil
}
