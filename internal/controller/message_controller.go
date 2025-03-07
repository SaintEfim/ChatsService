package controller

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/validator"

	"github.com/google/uuid"
)

type MessageController struct {
	messageValidator *validator.MessageValidator
	rep              interfaces.Repository[entity.Message]
}

func NewMessageController(messageValidator *validator.MessageValidator, rep interfaces.Repository[entity.Message]) interfaces.MessageController {
	return &MessageController{messageValidator: messageValidator, rep: rep}
}

func (c *MessageController) Get(ctx context.Context) ([]*dto.Message, error) {
	messages := make([]*dto.Message, 0)

	messagesEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, messageEntity := range messagesEntities {
		message := &dto.Message{
			Id:         messageEntity.Id,
			ChatId:     messageEntity.ChatId,
			SenderId:   messageEntity.SenderId,
			ReceiverId: messageEntity.ReceiverId,
			Text:       messageEntity.Text,
			CreatedAt:  messageEntity.CreatedAt,
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (c *MessageController) GetMessagesByChatId(ctx context.Context, chatId uuid.UUID) ([]*dto.Message, error) {
	messages := make([]*dto.Message, 0)

	messagesEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, messageEntity := range messagesEntities {
		if messageEntity.ChatId == chatId {
			message := &dto.Message{
				Id:         messageEntity.Id,
				ChatId:     messageEntity.ChatId,
				SenderId:   messageEntity.SenderId,
				ReceiverId: messageEntity.ReceiverId,
				Text:       messageEntity.Text,
				CreatedAt:  messageEntity.CreatedAt,
			}

			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	messageEntity, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	message := &dto.Message{
		Id:         messageEntity.Id,
		ChatId:     messageEntity.ChatId,
		SenderId:   messageEntity.SenderId,
		ReceiverId: messageEntity.ReceiverId,
		Text:       messageEntity.Text,
		CreatedAt:  messageEntity.CreatedAt,
	}

	return message, nil
}

func (c *MessageController) Create(ctx context.Context, message *dto.MessageCreate) (*dto.Message, error) {
	if err := c.messageValidator.Validate(message); err != nil {
		return nil, err
	}

	createRes, err := c.rep.Create(ctx, &entity.Message{
		ChatId:     message.ChatId,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		Text:       message.Text,
	})
	if err != nil {
		return nil, err
	}

	createItem := &dto.Message{
		Id:         createRes.Id,
		ChatId:     message.ChatId,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		Text:       message.Text,
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

func (c *MessageController) Update(ctx context.Context, id uuid.UUID, message *dto.MessageUpdate) error {
	err := c.rep.Update(ctx, id, &entity.Message{
		Text: message.Text,
	})
	if err != nil {
		return err
	}

	return nil
}
