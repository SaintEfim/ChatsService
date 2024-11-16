package controller

import (
	"context"

	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/models/model"

	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
)

type MessageController struct {
	rep interfaces.Repository[entity.MessageEntity]
}

func NewMessageController(rep interfaces.Repository[entity.MessageEntity]) interfaces.Controller[model.MessageModel] {
	return &MessageController{rep: rep}
}

func (c *MessageController) Get(ctx context.Context) ([]*model.MessageModel, error) {
	messageModels := make([]*model.MessageModel, 0)
	messages, err := c.rep.Get(ctx)

	if err != nil {
		return nil, err
	}

	if err := deepcopier.Copy(messages).To(messageModels); err != nil {
		return nil, err
	}

	return messageModels, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*model.MessageModel, error) {
	messageModel := &model.MessageModel{}
	message, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	if err := deepcopier.Copy(message).To(messageModel); err != nil {
		return nil, err
	}

	return messageModel, nil
}

func (c *MessageController) Create(ctx context.Context, message *model.MessageModel) (uuid.UUID, error) {
	messageEntity := &entity.MessageEntity{}

	if err := deepcopier.Copy(message).To(messageEntity); err != nil {
		return uuid.Nil, err
	}

	createItemId, err := c.rep.Create(ctx, messageEntity)
	if err != nil {
		return uuid.Nil, err
	}

	return createItemId, err
}

func (c *MessageController) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *MessageController) Update(ctx context.Context, id uuid.UUID, message *model.MessageModel) error {
	messageEntity := &entity.MessageEntity{}

	if err := deepcopier.Copy(message).To(messageEntity); err != nil {
		return err
	}

	err := c.rep.Update(ctx, id, messageEntity)
	if err != nil {
		return err
	}

	return nil
}
