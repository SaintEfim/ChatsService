package controller

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type ChatController struct {
	rep interfaces.Repository[dto.Chat, dto.ChatDetail, dto.ChatCreate, dto.ChatUpdate]
}

func NewChatController(rep interfaces.Repository[dto.Chat, dto.ChatDetail, dto.ChatCreate, dto.ChatUpdate]) interfaces.Controller[dto.Chat, dto.ChatDetail, dto.ChatCreate, dto.ChatUpdate] {
	return &ChatController{rep: rep}
}

func (c *ChatController) Get(ctx context.Context) ([]*dto.Chat, error) {
	chats, err := c.rep.Get(ctx)

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (c *ChatController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error) {
	chat, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *ChatController) Create(ctx context.Context, chat *dto.ChatCreate) (*dto.ChatDetail, error) {
	createItem, err := c.rep.Create(ctx, chat)
	if err != nil {
		return nil, err
	}

	return createItem, nil
}

func (c *ChatController) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatController) Update(ctx context.Context, id uuid.UUID, chat *dto.ChatUpdate) error {
	err := c.rep.Update(ctx, id, chat)
	if err != nil {
		return err
	}

	return nil
}
