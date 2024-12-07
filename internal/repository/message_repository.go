package repository

import (
	"context"
	"database/sql"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	db    *sqlx.DB
	query interfaces.Query[dto.Message]
}

func NewMessageRepository(db *sqlx.DB, query interfaces.Query[dto.Message]) interfaces.Repository[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate] {
	return &MessageRepository{
		db:    db,
		query: query}
}

func (r *MessageRepository) Get(ctx context.Context) ([]*dto.Message, error) {
	messages := make([]*dto.Message, 0)

	err := r.db.SelectContext(ctx, &messages, r.query.Get())
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	message := &dto.Message{}

	if err := r.db.GetContext(ctx, &message, r.query.GetOneById(), id); err != nil {
		return nil, err
	}

	if message == nil {
		return nil, fmt.Errorf("message not found")
	}

	return message, nil
}

func (r *MessageRepository) Create(ctx context.Context, message *dto.MessageCreate) (uuid.UUID, error) {
	id := uuid.New()

	_, err := r.db.ExecContext(ctx, r.query.Create(),
		id,
		message.ChatId,
		message.EmployeeId,
		message.Text)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.GetOneById(ctx, id)
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, r.query.Delete(), id)
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

	result, err := r.db.ExecContext(ctx, r.query.Update(),
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
