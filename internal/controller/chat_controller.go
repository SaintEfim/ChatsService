package controller

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/google/uuid"
)

type ChatController struct {
	rep interfaces.Repository[entity.Chat]
}

func NewChatController(rep interfaces.Repository[entity.Chat]) interfaces.Controller[dto.Chat, dto.ChatDetail, dto.ChatCreate, dto.ChatUpdate] {
	return &ChatController{rep: rep}
}

func (c *ChatController) Get(ctx context.Context) ([]*dto.Chat, error) {
	var chats []*dto.Chat

	chatEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, chatEntity := range chatEntities {
		chat := &dto.Chat{
			Id:   chatEntity.Id,
			Name: chatEntity.Name,
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (c *ChatController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error) {
	chatEntity, err := c.rep.GetOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	chat := &dto.ChatDetail{
		Id:          chatEntity.Id,
		Name:        chatEntity.Name,
		IsGroup:     chatEntity.IsGroup,
		EmployeeIds: chatEntity.EmployeeIds,
	}

	return chat, nil
}

func (c *ChatController) Create(ctx context.Context, chat *dto.ChatCreate) (*dto.ChatDetail, error) {
	createRes, err := c.rep.Create(ctx, &entity.Chat{
		Name:        chat.Name,
		IsGroup:     chat.IsGroup,
		EmployeeIds: chat.EmployeeIds,
	})
	if err != nil {
		return nil, err
	}

	createItem := &dto.ChatDetail{
		Id:          createRes.Id,
		Name:        createRes.Name,
		IsGroup:     createRes.IsGroup,
		EmployeeIds: createRes.EmployeeIds,
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
	err := c.rep.Update(ctx, id, &entity.Chat{
		Name:        chat.Name,
		EmployeeIds: chat.EmployeeIds,
	})
	if err != nil {
		return err
	}

	return nil
}
