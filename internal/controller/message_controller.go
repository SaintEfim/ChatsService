package controller

import (
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/models/model"
	"context"

	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
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

	automapper.MapLoose(messages, &messageModels)

	return messageModels, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*model.MessageModel, error) {
	messageModel := &model.MessageModel{}
	message, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	automapper.MapLoose(message, messageModel)

	return messageModel, nil
}

func (c *MessageController) Create(ctx context.Context, message *model.MessageModel) (uuid.UUID, error) {
	messageEntity := &entity.MessageEntity{}

	automapper.MapLoose(message, messageEntity)

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

	automapper.MapLoose(message, messageEntity)

	err := c.rep.Update(ctx, id, messageEntity)
	if err != nil {
		return err
	}

	return nil
}
