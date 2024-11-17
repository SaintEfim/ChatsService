package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	retrieveAllChats = `SELECT id, name, is_group, employee_ids FROM chats`
	retrieveChatById = `SELECT id, name, is_group, employee_ids FROM chats WHERE id = $1`
	createChat       = `INSERT INTO chats (id, name, is_group, employee_ids, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	deleteChat       = `DELETE FROM chats WHERE id = $1`
	updateChat       = `UPDATE chats SET name = $1, employee_ids = $2, updated_at = NOW() WHERE id = $3`
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) interfaces.Repository[entity.ChatEntity] {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*entity.ChatEntity, error) {
	chats := make([]*entity.ChatEntity, 0)

	rows, err := r.db.QueryxContext(ctx, retrieveAllChats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat entity.ChatEntity
		var rawEmployeeIDs []byte

		err := rows.Scan(&chat.Id, &chat.Name, &chat.IsGroup, &rawEmployeeIDs)
		if err != nil {
			return nil, err
		}

		if len(rawEmployeeIDs) > 0 {
			var employeeIds []string
			err = json.Unmarshal(rawEmployeeIDs, &employeeIds)
			if err != nil {
				return nil, err
			}

			for _, empId := range employeeIds {
				uuidValue, err := uuid.Parse(empId)
				if err != nil {
					return nil, err
				}
				chat.EmployeeIds = append(chat.EmployeeIds, uuidValue)
			}
		} else {
			chat.EmployeeIds = make([]uuid.UUID, 0)
		}

		chats = append(chats, &chat)
	}

	if err = rows.Err(); err != nil {
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

func (r *ChatRepository) Create(ctx context.Context, chat *entity.ChatEntity) (uuid.UUID, error) {
	chat.Id = uuid.New()

	_, err := r.db.ExecContext(ctx, createChat,
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
	result, err := r.db.ExecContext(ctx, updateChat,
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
