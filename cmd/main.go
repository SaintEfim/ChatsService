package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	"ChatsService/config"
	"ChatsService/internal/controller"
	"ChatsService/internal/database/psql"
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
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		fx.Provide(func() (*config.Config, error) {
			return config.ReadConfig("config", "yaml", "./config")
		}),
		fx.Invoke(func(ctx context.Context, cfg *config.Config, logger *zap.Logger) error {
			return psql.CreateDatabase(ctx, cfg, logger)
		}),
		fx.Provide(
			logger.NewLogger,
			repository.NewChatRepository,
			repository.NewMessageRepository,
			repository.NewEmployeeChatSettingsRepository,
			controller.NewChatController,
			psql.PostgresConnect,
			server.NewHTTPServer,
			server.NewServer,
			handler.NewChatHandler,
			handler.NewMessageHandler,
		),
		fx.Invoke(registerServer),
		fx.Invoke(registerPostgres),
	)
}
