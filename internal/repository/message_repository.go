package repository

import (
	"context"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) interfaces.Repository[entity.Message] {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Get(ctx context.Context) ([]*entity.Message, error) {
	messages := make([]*entity.Message, 0)

	if err := r.db.WithContext(ctx).Model(&entity.Message{}).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Message, error) {
	message := &entity.Message{}

	if err := r.db.WithContext(ctx).Model(&entity.Message{}).First(&message, id).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, entityCreate *entity.Message) (*entity.Message, error) {
	if err := r.db.WithContext(ctx).Model(&entity.Message{}).Create(&entityCreate).Error; err != nil {
		return nil, err
	}

	return entityCreate, nil
}

func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&entity.Message{}).Delete(&entity.Message{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *MessageRepository) Update(ctx context.Context, id uuid.UUID, updates *entity.Message) error {
	entityUpdates := &entity.Chat{}
	if err := r.db.WithContext(ctx).Model(&entity.Chat{}).First(&entityUpdates, id).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).
		Model(entityUpdates).
		Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
