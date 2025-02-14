package postgres

import (
	"github.com/only1nft/gormzap"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ChatsService/config"
	"ChatsService/internal/models/entity"
)

func ConnectToDB(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DataBase.ConnectionString), &gorm.Config{Logger: gormzap.New(logger)})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Chat{}, &entity.Message{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
