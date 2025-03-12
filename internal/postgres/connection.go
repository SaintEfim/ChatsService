package postgres

import (
	"database/sql"

	"ChatsService/config"
	"ChatsService/internal/models/entity"

	_ "github.com/lib/pq"
	"github.com/only1nft/gormzap"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	fallbackDB, err := sql.Open("postgres", cfg.DataBase.ConnectionPostgres)
	if err != nil {
		logger.Error("failed to connect to fallback database", zap.Error(err))
		return nil, err
	}
	defer fallbackDB.Close()

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err = fallbackDB.QueryRow(query, cfg.DataBase.Name).Scan(&exists)
	if err != nil {
		logger.Error("failed to check database existence", zap.Error(err))
		return nil, err
	}

	if !exists {
		_, err = fallbackDB.Exec("CREATE DATABASE " + cfg.DataBase.Name)
		if err != nil {
			logger.Error("failed to create database", zap.String("database", cfg.DataBase.Name), zap.Error(err))
			return nil, err
		}
		logger.Info("Database created", zap.String("database", cfg.DataBase.Name))
	} else {
		logger.Info("Database already exists", zap.String("database", cfg.DataBase.Name))
	}

	db, err := gorm.Open(postgres.Open(cfg.DataBase.ConnectionString), &gorm.Config{
		Logger: gormzap.New(logger),
	})
	if err != nil {
		logger.Error("failed to connect to target database", zap.Error(err))
		return nil, err
	}

	err = db.AutoMigrate(&entity.Chat{}, &entity.Message{})
	if err != nil {
		logger.Error("failed to migrate database", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to database", zap.String("database", cfg.DataBase.Name))
	return db, nil
}
