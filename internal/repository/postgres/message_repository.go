package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	getAllMessages = `SELECT id, chat_id, employee_id, colleague_id, text, created_at FROM messages`
	getMessageById = `SELECT id, chat_id, employee_id, colleague_id, text FROM messages WHERE id = $1`
	createMessage  = `INSERT INTO messages (id, chat_id, employeeId, colleague_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteMessage  = `DELETE FROM messages WHERE id = $1`
	updateMessage  = `UPDATE messages SET text = $1, updated_at = NOW() WHERE id = $2`
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) interfaces.Repository[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate] {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Get(ctx context.Context) ([]*dto.Message, error) {
	messages := make([]*dto.Message, 0)

	err := r.db.SelectContext(ctx, &messages, getAllMessages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	message := &dto.Message{}

	if err := r.db.GetContext(ctx, &message, getMessageById, id); err != nil {
		return nil, err
	}

	if message == nil {
		return nil, fmt.Errorf("message not found")
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, message *dto.MessageCreate) (*dto.Message, error) {
	id := uuid.New()

	_, err := r.db.ExecContext(ctx, createMessage,
		id,
		message.ChatId,
		message.EmployeeId,
		message.Text)
	if err != nil {
		return nil, err
	}

	createItem, err := r.GetOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	return createItem, nil
}

func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.GetOneById(ctx, id)
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, deleteMessage, id)
	if err != nil {
		return err
	}

	if err := r.checkRows(ctx, result); err != nil {
		return err
	}

	return nil
}

func (r *MessageRepository) Update(ctx context.Context, id uuid.UUID, message *dto.MessageUpdate) error {
	_, err := r.GetOneById(ctx, id)
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, updateMessage,
		message.Text,
		id)
	if err != nil {
		return err
	}

	if err := r.checkRows(ctx, result); err != nil {
		return err
	}

	return nil
}

func (r *MessageRepository) checkRows(ctx context.Context, result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
