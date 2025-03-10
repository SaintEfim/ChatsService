package interfaces

import (
	"ChatsService/proto/chat"
	"context"
)

type ChatGRPC interface {
	CreateMessage(ctx context.Context, req *chat.MessageCreateRequest) (*chat.MessageCreateResponse, error)
	chat.GreeterChatsServer
}
