package psql

import (
	"context"
	"fmt"

	"ChatsService/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	checkDb               = `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`
	createDbFormat        = `CREATE DATABASE "%s"`
	adminConnectionString = "postgres://postgres:1234@localhost/postgres?sslmode=disable"
)

func CreateDatabase(ctx context.Context, cfg *config.Config, logger *zap.Logger) error {
	var exists bool

	adminDB, err := sqlx.Open(cfg.DataBase.DriverName, adminConnectionString)
	if err != nil {
		return err
	}

	err = adminDB.QueryRowContext(ctx, checkDb, cfg.DataBase.DataBaseName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		createDbQuery := fmt.Sprintf(createDbFormat, cfg.DataBase.DataBaseName)
		_, err = adminDB.Exec(createDbQuery)
		if err != nil {
			return err
		}

		logger.Sugar().Infof("Database %s successfully created", cfg.DataBase.DataBaseName)
	} else {
		logger.Sugar().Infof("Database %s already exists", cfg.DataBase.DataBaseName)
	}

	return nil
}
