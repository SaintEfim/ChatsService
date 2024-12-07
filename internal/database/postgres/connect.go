package postgres

import (
	"ChatsService/config"

	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.DataBase.ConnectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
