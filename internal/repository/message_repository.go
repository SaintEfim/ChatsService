package repository

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	retrieveAllMessages = `SELECT id, chat_id, employee_id, text FROM messages`
	retrieveMessageById = `SELECT id, chat_id, employee_id, text FROM messages WHERE id = $1`
	createMessage       = `INSERT INTO messages (id, chat_id, employee_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteMessage       = `DELETE FROM messages WHERE id = $1`
	updateMessage       = `UPDATE messages SET text = $1, updated_at = NOW() WHERE id = $2`
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) interfaces.Repository[entity.MessageEntity] {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Get(ctx context.Context) ([]*entity.MessageEntity, error) {
	messages := make([]*entity.MessageEntity, 0)

	err := r.db.SelectContext(ctx, &messages, retrieveAllMessages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.MessageEntity, error) {
	message := &entity.MessageEntity{}

	if err := r.db.GetContext(ctx, &message, retrieveMessageById, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, message *entity.MessageEntity) error {
	message.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createMessage,
		message.Id,
		message.ChatId,
		message.EmployeeId,
		message.Text)
	if err != nil {
		return err
	}

	return err
}

func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteMessage, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (r *MessageRepository) Update(ctx context.Context, id uuid.UUID, message *entity.MessageEntity) error {
	result, err := r.db.ExecContext(ctx, updateMessage,
		message.Text,
		id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
