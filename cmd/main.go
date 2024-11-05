package cmd

import (
	"context"

	"ChatsService/config"
	"ChatsService/internal/database"
	"ChatsService/internal/repository/chat"
	"ChatsService/internal/repository/message"
	"ChatsService/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func registerPostgres(lc fx.Lifecycle, db *sqlx.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.PingContext(ctx); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}

func main() {
	fx.New(
		fx.Provide(func() (*config.Config, error) {
			return config.ReadConfig("config", "yaml", "./config")
		}),
		fx.Provide(
			logger.NewLogger,
			chat.NewChatRepository,
			message.NewMessageRepository,
			database.PostgresConnect,
		),
		fx.Invoke(registerPostgres),
	)
}
