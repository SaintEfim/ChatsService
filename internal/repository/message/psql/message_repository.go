package psql

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) interfaces.MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) GetAllByChatId(ctx context.Context, chatId uuid.UUID) ([]*entity.MessageEntity, error) {
	var messages []*entity.MessageEntity

	rows, err := r.db.QueryContext(ctx, retrieveAllMessagesByChatId, chatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := &entity.MessageEntity{}
		if err := rows.Scan(&message.Id, &message.ChatId, &message.EmployeeId, &message.Text); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.MessageEntity, error) {
	message := &entity.MessageEntity{}

	if err := r.db.QueryRowContext(ctx, retrieveMessageById, id).Scan(&message.Id, &message.EmployeeId, &message.Text); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, message *entity.MessageEntity) error {
	message.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createMessage, message.Id, message.ChatId, message.EmployeeId, message.Text)
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
	result, err := r.db.ExecContext(ctx, updateMessage, message.Text, id)

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
