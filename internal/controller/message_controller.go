package controller

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type MessageController struct {
	rep interfaces.Repository[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate]
}

func NewMessageController(rep interfaces.Repository[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate]) interfaces.Controller[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate] {
	return &MessageController{rep: rep}
}

func (c *MessageController) Get(ctx context.Context) ([]*dto.Message, error) {
	messages, err := c.rep.Get(ctx)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	message, err := c.rep.GetOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (c *MessageController) Create(ctx context.Context, message *dto.MessageCreate) (*dto.Message, error) {
	createItem, err := c.rep.Create(ctx, message)
	if err != nil {
		return nil, err
	}

	return createItem, err
}

func (c *MessageController) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *MessageController) Update(ctx context.Context, id uuid.UUID, message *dto.MessageUpdate) error {
	err := c.rep.Update(ctx, id, message)
	if err != nil {
		return err
	}

	return nil
}
