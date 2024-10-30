package connection

import (
	"ChatsService/config"

	"github.com/jmoiron/sqlx"
)

func PostgresConnect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionStrings.ServiceDb)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
