package controller

import (
	"context"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/models/model"

	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
)

type ChatController struct {
	rep interfaces.Repository[entity.ChatEntity]
}

func NewChatController(rep interfaces.Repository[entity.ChatEntity]) interfaces.Controller[model.ChatModel] {
	return &ChatController{rep: rep}
}

func (c *ChatController) Get(ctx context.Context) ([]*model.ChatModel, error) {
	chatModels := make([]*model.ChatModel, 0)
	chats, err := c.rep.Get(ctx)

	if err != nil {
		return nil, err
	}

	if err := deepcopier.Copy(chats).To(chatModels); err != nil {
		return nil, err
	}

	return chatModels, nil
}

func (c *ChatController) GetOneById(ctx context.Context, id uuid.UUID) (*model.ChatModel, error) {
	chatModel := &model.ChatModel{}
	chat, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	if err := deepcopier.Copy(chat).To(chatModel); err != nil {
		return nil, err
	}

	return chatModel, nil
}

func (c *ChatController) Create(ctx context.Context, chat *model.ChatModel) (uuid.UUID, error) {
	chatEntity := &entity.ChatEntity{}

	if err := deepcopier.Copy(chat).To(chatEntity); err != nil {
		return uuid.Nil, err
	}

	createItemId, err := c.rep.Create(ctx, chatEntity)
	if err != nil {
		return uuid.Nil, err
	}

	return createItemId, nil
}

func (c *ChatController) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatController) Update(ctx context.Context, id uuid.UUID, chat *model.ChatModel) error {
	chatEntity := &entity.ChatEntity{}

	if err := deepcopier.Copy(chat).To(chatEntity); err != nil {
		return err
	}

	err := c.rep.Update(ctx, id, chatEntity)
	if err != nil {
		return err
	}

	return nil
}
