package grpc

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/chat"

	"github.com/google/uuid"
)

type ChatGRPCServer struct {
	controller interfaces.MessageController
	chat.UnimplementedGreeterChatsServer
}

func NewChatGrpcServer(controller interfaces.MessageController) interfaces.ChatGRPCServer {
	return &ChatGRPCServer{controller: controller}
}

func (c *ChatGRPCServer) CreateMessage(ctx context.Context, req *chat.MessageCreateRequest) (*chat.MessageCreateResponse, error) {
	response := &chat.MessageCreateResponse{}

	request, err := c.createRequestModel(ctx, req)
	if err != nil {
		return nil, err
	}

	message, err := c.controller.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	response.Id = message.Id.String()
	response.CreateAt = message.CreatedAt.String()

	return response, nil
}

func (c *ChatGRPCServer) createRequestModel(ctx context.Context, req *chat.MessageCreateRequest) (*dto.MessageCreate, error) {
	ChatId, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, err
	}

	SenderId, err := uuid.Parse(req.SenderId)
	if err != nil {
		return nil, err
	}

	ReceiverId, err := uuid.Parse(req.ReceiverId)
	if err != nil {
		return nil, err
	}

	request := &dto.MessageCreate{
		ChatId:     ChatId,
		SenderId:   SenderId,
		ReceiverId: ReceiverId,
		Text:       req.Text,
	}

	return request, nil
}
