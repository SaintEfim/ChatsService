package chat

import (
	"context"
	"database/sql"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) interfaces.Repository[entity.ChatEntity] {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*entity.ChatEntity, error) {
	var chats []*entity.ChatEntity

	rows, err := r.db.QueryContext(ctx, retrieveAllChats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := &entity.ChatEntity{}
		if err := rows.Scan(&chat.Id, &chat.Name, pq.Array(&chat.EmployeesIds)); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) GetOneById(ctx context.Context, id uuid.UUID) (*entity.ChatEntity, error) {
	chat := &entity.ChatEntity{}

	err := r.db.QueryRowContext(ctx, retrieveChatById, id).Scan(&chat.Id, &chat.Name, pq.Array(&chat.EmployeesIds))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) Create(ctx context.Context, chat *entity.ChatEntity) error {
	chat.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createChat, chat.Id, chat.Name, pq.Array(chat.EmployeesIds))
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
	result, err := r.db.ExecContext(ctx, updateChat, chat.Name, pq.Array(chat.EmployeesIds), id)
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
