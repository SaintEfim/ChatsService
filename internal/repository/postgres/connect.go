package postgres

import (
	"database/sql"

	"ChatsService/config"
	"ChatsService/internal/models/entity"

	"github.com/only1nft/gormzap"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config, logger *zap.Logger) (*gorm.DB, *sql.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DataBase.ConnectionString), &gorm.Config{Logger: gormzap.New(logger)})
	if err != nil {
		return nil, nil, err
	}

	err = db.AutoMigrate(&entity.Chat{}, &entity.Message{})
	if err != nil {
		return nil, nil, err
	}

	psql, err := db.DB()
	if err != nil {
		return nil, psql, err
	}

	return db, psql, nil
}
