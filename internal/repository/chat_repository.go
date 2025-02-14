package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) interfaces.Repository[entity.Chat] {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*entity.Chat, error) {
	chats := make([]*entity.Chat, 0)

	if err := r.db.WithContext(ctx).Model(&entity.Chat{}).Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Chat, error) {
	chat := &entity.Chat{}

	if err := r.db.WithContext(ctx).Model(&entity.Chat{}).First(&chat, id).Error; err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) Create(ctx context.Context, entityCreate *entity.Chat) (*entity.Chat, error) {
	if err := r.db.WithContext(ctx).Model(&entity.Chat{}).Create(&entityCreate).Error; err != nil {
		return nil, err
	}

	return entityCreate, nil
}

func (r *ChatRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&entity.Chat{}).Delete(&entity.Chat{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) Update(ctx context.Context, id uuid.UUID, updates *entity.Chat) error {
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
