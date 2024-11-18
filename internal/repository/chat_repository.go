package repository

import (
	"context"
	"fmt"

	"ChatsService/internal/database/postgres/query"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatRepository struct {
	db interfaces.QueryExecutor
}

func NewChatRepository(db interfaces.QueryExecutor) interfaces.Repository[entity.ChatEntity] {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*entity.ChatEntity, error) {
	chats := make([]*entity.ChatEntity, 0)

	rows, err := r.db.QueryContext(ctx, query.GetAllChats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := &entity.ChatEntity{}

		err := rows.Scan(
			&chat.Id,
			&chat.Name,
			&chat.IsGroup,
			pq.Array(&chat.EmployeeIds),
		)
		if err != nil {
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

	rows, err := r.db.QueryContext(ctx, query.GetChatById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&chat.Id,
			&chat.Name,
			&chat.IsGroup,
			pq.Array(&chat.EmployeeIds),
		)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("chat with id %s not found", id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) Create(ctx context.Context, chat *entity.ChatEntity) (uuid.UUID, error) {
	chat.Id = uuid.New()

	if chat.EmployeeIds == nil {
		chat.EmployeeIds = make([]uuid.UUID, 0)
	}

	_, err := r.db.ExecContext(ctx, query.CreateChat,
		chat.Id,
		chat.Name,
		chat.IsGroup,
		pq.Array(chat.EmployeeIds))
	if err != nil {
		return uuid.Nil, err
	}

	return chat.Id, nil
}

func (r *ChatRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, query.DeleteChat, id)
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
	if chat.EmployeeIds == nil {
		chat.EmployeeIds = make([]uuid.UUID, 0)
	}

	result, err := r.db.ExecContext(ctx, query.UpdateChat,
		chat.Name,
		pq.Array(chat.EmployeeIds),
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
