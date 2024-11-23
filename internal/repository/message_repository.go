package repository

import (
	"ChatsService/internal/exception"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type MessageRepository struct {
	db    interfaces.QueryExecutor
	query interfaces.Query[entity.MessageEntity]
}

func NewMessageRepository(db interfaces.QueryExecutor, query interfaces.Query[entity.MessageEntity]) interfaces.Repository[entity.MessageEntity] {
	return &MessageRepository{
		db:    db,
		query: query}
}

func (r *MessageRepository) Get(ctx context.Context) ([]*entity.MessageEntity, error) {
	messages := make([]*entity.MessageEntity, 0)

	err := r.db.SelectContext(ctx, &messages, r.query.Get())
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.MessageEntity, error) {
	message := &entity.MessageEntity{}

	if err := r.db.GetContext(ctx, &message, r.query.GetOneById(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if message == nil {
		return nil, exception.NewNotFoundException(fmt.Sprintf("Message with id %s not found", id))
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, message *entity.MessageEntity) (uuid.UUID, error) {
	message.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, r.query.Create(),
		message.Id,
		message.ChatId,
		message.EmployeeId,
		message.Text)
	if err != nil {
		return uuid.Nil, err
	}

	return message.Id, nil
}

func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, r.query.Delete(), id)
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
	result, err := r.db.ExecContext(ctx, r.query.Update(),
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
