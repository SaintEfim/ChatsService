package cmd

import (
	"context"
	"fmt"

	"ChatsService/config"
	"ChatsService/internal/database"
	"ChatsService/internal/handler"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/repository"
	"ChatsService/internal/server"
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

func registerServer(ctx context.Context, lifecycle fx.Lifecycle, srv interfaces.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			go func() {
				err = srv.Run(ctx)
			}()

			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Stop(ctx); err != nil {
				return fmt.Errorf("failed to stop server: %w", err)
			}
			return nil
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
			repository.NewChatRepository,
			repository.NewMessageRepository,
			database.PostgresConnect,
			server.NewHTTPServer,
			server.NewServer,
			handler.NewChatHandler,
			handler.NewMessageHandler,
		),
		fx.Invoke(registerPostgres),
	)
}
