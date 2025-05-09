package postgres

import (
	"ChatsService/config"
	"ChatsService/internal/models/entity"

	_ "github.com/lib/pq"
	"github.com/only1nft/gormzap"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
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
