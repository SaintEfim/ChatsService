package controller

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type MessageController struct {
	rep interfaces.Repository[entity.Message]
}

func NewMessageController(rep interfaces.Repository[entity.Message]) interfaces.Controller[dto.Message, dto.Message, dto.MessageCreate, dto.MessageUpdate] {
	return &MessageController{rep: rep}
}

func (c *MessageController) Get(ctx context.Context) ([]*dto.Message, error) {
	var chatDtos []*dto.Message

	chatEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, chatEntity := range chatEntities {
		chatDto := &dto.Message{
			Id:          chatEntity.Id,
			ChatId:      chatEntity.ChatId,
			EmployeeId:  chatEntity.EmployeeId,
			ColleagueId: chatEntity.ColleagueId,
			Text:        chatEntity.Text,
			CreatedAt:   chatEntity.CreatedAt,
		}

		chatDtos = append(chatDtos, chatDto)
	}

	return chatDtos, nil
}

func (c *MessageController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	chatEntity, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	chatDto := &dto.Message{
		Id:          chatEntity.Id,
		ChatId:      chatEntity.ChatId,
		EmployeeId:  chatEntity.EmployeeId,
		ColleagueId: chatEntity.ColleagueId,
		Text:        chatEntity.Text,
		CreatedAt:   chatEntity.CreatedAt,
	}

	return chatDto, nil
}

func (c *MessageController) Create(ctx context.Context, chat *dto.MessageCreate) (*dto.Message, error) {
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

func (c *MessageController) Update(ctx context.Context, id uuid.UUID, chat *dto.MessageUpdate) error {
	err := c.rep.Update(ctx, id, &entity.Message{
		Text: chat.Text,
	})
	if err != nil {
		return err
	}

	return nil
}
