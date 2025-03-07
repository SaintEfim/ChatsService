package controller

import (
	"context"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/validator"
	"ChatsService/proto/employee"

	"github.com/google/uuid"
)

type ChatController struct {
	chatValidator  *validator.ChatValidator
	rep            interfaces.Repository[entity.Chat]
	employeeClient interfaces.EmployeeGrpc
}

func NewChatController(chatValidator *validator.ChatValidator,
	rep interfaces.Repository[entity.Chat],
	employeeClient interfaces.EmployeeGrpc) interfaces.ChatController {
	return &ChatController{chatValidator: chatValidator, rep: rep, employeeClient: employeeClient}
}

func (c *ChatController) Get(ctx context.Context) ([]*dto.Chat, error) {
	chats := make([]*dto.Chat, 0)

	chatEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, chatEntity := range chatEntities {
		participant, err := c.fetchEmployees(ctx, chatEntity.ParticipantIds)
		if err != nil {
			return nil, err
		}

		chat := &dto.Chat{
			Id:           chatEntity.Id,
			Name:         chatEntity.Name,
			Participants: participant,
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (c *ChatController) GetChatsByUserId(ctx context.Context, userId uuid.UUID) ([]*dto.Chat, error) {
	chats := make([]*dto.Chat, 0)

	chatEntities, err := c.rep.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, chatEntity := range chatEntities {
		if !chatEntity.ParticipantIds.Contains(userId) {
			continue
		}

		participants, err := c.fetchEmployees(ctx, chatEntity.ParticipantIds)
		if err != nil {
			return nil, err
		}

		chat := &dto.Chat{
			Id:           chatEntity.Id,
			Name:         chatEntity.Name,
			Participants: participants,
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (c *ChatController) PrivateChatExists(ctx context.Context, userId uuid.UUID, colleagueId uuid.UUID) (bool, error) {
	if userId == colleagueId {
		return false, nil
	}

	chatEntities, err := c.rep.Get(ctx)
	if err != nil {
		return false, err
	}

	for _, chatEntity := range chatEntities {
		if chatEntity.ParticipantIds.Contains(userId) && chatEntity.ParticipantIds.Contains(colleagueId) {
			return true, nil
		}
	}

	return false, nil
}

func (c *ChatController) GetOneById(ctx context.Context, id uuid.UUID) (*dto.ChatDetail, error) {
	chatEntity, err := c.rep.GetOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	participants, err := c.fetchEmployees(ctx, chatEntity.ParticipantIds)
	if err != nil {
		return nil, err
	}

	chat := &dto.ChatDetail{
		Id:           chatEntity.Id,
		Name:         chatEntity.Name,
		IsGroup:      chatEntity.IsGroup,
		Participants: participants,
	}

	return chat, nil
}

func (c *ChatController) Create(ctx context.Context, chat *dto.ChatCreate) (*dto.CreateAction, error) {
	if err := c.chatValidator.Validate(chat); err != nil {
		return nil, err
	}

	createRes, err := c.rep.Create(ctx, &entity.Chat{
		Name:           chat.Name,
		IsGroup:        chat.IsGroup,
		ParticipantIds: chat.ParticipantIds,
	})
	if err != nil {
		return nil, err
	}

	createAction := &dto.CreateAction{
		Id: createRes.Id,
	}

	return createAction, nil
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
		Name:           chat.Name,
		ParticipantIds: chat.ParticipantIds,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatController) fetchEmployees(ctx context.Context, employeeIDs []uuid.UUID) ([]dto.Participant, error) {
	ids := entity.UUIDArray(employeeIDs).ToStringSlice()

	employeesResponse, err := c.employeeClient.Search(ctx, &employee.SearchRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employees: %w", err)
	}

	employees := make([]dto.Participant, len(employeesResponse.Employees))
	for i, emp := range employeesResponse.Employees {
		employees[i] = dto.Participant{
			Id:         uuid.MustParse(emp.Id),
			Name:       emp.Name,
			Surname:    emp.Surname,
			Patronymic: emp.Patronymic,
		}
	}

	return employees, nil
}
