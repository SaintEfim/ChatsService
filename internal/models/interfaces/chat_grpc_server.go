package interfaces

import (
	"context"

	"ChatsService/proto/chat"
)

type ChatGRPCServer interface {
	CreateMessage(ctx context.Context, req *chat.MessageCreateRequest) (*chat.MessageCreateResponse, error)
	chat.GreeterChatsServer
}
