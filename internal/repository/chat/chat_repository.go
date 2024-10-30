package chat

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ChatRepository struct {
	db *sqlx.DB
}

const (
	retrieveAllChats = `SELECT id, name, employees_ids FROM chats`
	retrieveChatById = `SELECT id, name, employees_ids FROM chats WHERE id = $1`
	createChat       = `INSERT INTO chats (id, name, employees_ids) VALUES ($1, $2, $3)`
	deleteChat       = `DELETE FROM chats WHERE id = $1`
	updateChat       = `UPDATE chats SET name = $1, employees_ids = $2 WHERE id = $3`
)

func NewChatRepository(db *sqlx.DB) interfaces.Repository[entity.ChatEntity] {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*entity.ChatEntity, error) {
	chats := make([]*entity.ChatEntity, 0)

	err := r.db.SelectContext(ctx, &chats, retrieveAllChats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.ChatEntity, error) {
	chat := &entity.ChatEntity{}

	if err := r.db.GetContext(ctx, &chat, retrieveChatById, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) Create(ctx context.Context, chat *entity.ChatEntity) error {
	chat.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createChat, chat.Id, chat.Name, pq.Array(chat.EmployeeIds))
	if err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteChat, id)
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

func (r *ChatRepository) Update(ctx context.Context, id uuid.UUID, chat *entity.ChatEntity) error {
	result, err := r.db.ExecContext(ctx, updateChat, chat.Name, pq.Array(chat.EmployeeIds), id)
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
