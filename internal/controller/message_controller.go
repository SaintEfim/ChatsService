package controller

import (
	"context"

	"github.com/google/uuid"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
)

type MessageController struct {
	rep interfaces.Repository[entity.Message]
}

func NewMessageController(rep interfaces.Repository[entity.Message]) interfaces.MessageController {
	return &MessageController{rep: rep}
}

func (c *MessageController) Get(ctx context.Context) ([]*dto.Message, error) {
	messages := make([]*dto.Message, 0)

	messagesEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, messageEntity := range messagesEntities {
		message := &dto.Message{
			Id:          messageEntity.Id,
			ChatId:      messageEntity.ChatId,
			EmployeeId:  messageEntity.EmployeeId,
			ColleagueId: messageEntity.ColleagueId,
			Text:        messageEntity.Text,
			CreatedAt:   messageEntity.CreatedAt,
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	messageEntity, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	message := &dto.Message{
		Id:          messageEntity.Id,
		ChatId:      messageEntity.ChatId,
		EmployeeId:  messageEntity.EmployeeId,
		ColleagueId: messageEntity.ColleagueId,
		Text:        messageEntity.Text,
		CreatedAt:   messageEntity.CreatedAt,
	}

	return message, nil
}

func (c *MessageController) Create(ctx context.Context, chat *dto.Message) (*dto.Message, error) {
	createRes, err := c.rep.Create(ctx, &entity.Message{
		ChatId:      chat.ChatId,
		EmployeeId:  chat.EmployeeId,
		ColleagueId: chat.ColleagueId,
		Text:        chat.Text,
	})
	if err != nil {
		return nil, err
	}

	createItem := &dto.Message{
		Id:          createRes.Id,
		ChatId:      chat.ChatId,
		EmployeeId:  chat.EmployeeId,
		ColleagueId: chat.ColleagueId,
		Text:        chat.Text,
	}

	return createItem, nil
}

func (c *MessageController) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *MessageController) Update(ctx context.Context, id uuid.UUID, chat *dto.Message) error {
	err := c.rep.Update(ctx, id, &entity.Message{
		Text: chat.Text,
	})
	if err != nil {
		return err
	}

	return nil
}
