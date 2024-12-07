package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatRepository struct {
	db    *sqlx.DB
	query interfaces.Query[dto.Chat]
}

func NewChatRepository(db *sqlx.DB, query interfaces.Query[dto.Chat]) interfaces.Repository[dto.Chat, dto.ChatDetail, dto.ChatCreate, dto.ChatUpdate] {
	return &ChatRepository{
		db:    db,
		query: query,
	}
}

func (r *ChatRepository) Get(ctx context.Context) ([]*dto.Chat, error) {
	chats := make([]*dto.Chat, 0)

	rows, err := r.db.QueryContext(ctx, r.query.Get())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := &dto.Chat{}

		err := rows.Scan(
			&chat.Id,
			&chat.Name,
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

func (r *ChatRepository) GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error) {
	chat := &dto.ChatDetail{}

	rows, err := r.db.QueryContext(ctx, r.query.GetOneById(), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = rows.Scan(
		&chat.Id,
		&chat.Name,
		&chat.IsGroup,
		pq.Array(&chat.EmployeeIds),
	)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) Create(ctx context.Context, entity *dto.ChatCreate) (uuid.UUID, error) {
	id := uuid.New()

	if entity.EmployeeIds == nil {
		entity.EmployeeIds = make([]uuid.UUID, 0)
	}

	_, err := r.db.ExecContext(ctx, r.query.Create(),
		id,
		entity.Name,
		entity.IsGroup,
		pq.Array(entity.EmployeeIds))
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *ChatRepository) Delete(ctx context.Context, id uuid.UUID) error {
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

func (r *ChatRepository) Update(ctx context.Context, id uuid.UUID, entity *dto.ChatUpdate) error {
	_, err := r.GetOneById(ctx, id)
	if err != nil {
		return err
	}

	if entity.EmployeeIds == nil {
		entity.EmployeeIds = make([]uuid.UUID, 0)
	}

	result, err := r.db.ExecContext(ctx, r.query.Update(),
		entity.Name,
		pq.Array(entity.EmployeeIds),
		id)
	if err != nil {
		return err
	}

	if err := r.checkRows(ctx, result); err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) checkRows(ctx context.Context, result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
