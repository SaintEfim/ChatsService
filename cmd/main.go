package main

import (
	"ChatsService/internal/validation"
	"context"

	"ChatsService/config"
	"ChatsService/internal/controller"
	"ChatsService/internal/delivery/grpc"
	"ChatsService/internal/handler"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/postgres"
	"ChatsService/internal/repository"
	"ChatsService/internal/server"
	"ChatsService/pkg/logger"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func registerPostgres(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}

			if err := sqlDB.PingContext(ctx); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}

			return sqlDB.Close()
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
				return err
			}
			return nil
		},
	})
}

func registerGRPCClient(lc fx.Lifecycle, client interfaces.EmployeeGrpc) {
	lc.Append(fx.Hook{
		OnStart: client.Initialize,
		OnStop:  client.Close,
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
		fx.Provide(
			logger.NewLogger,
			postgres.ConnectToDB,
			grpc.NewEmployeeGrpcClient,
			validation.NewChatValidator,
			repository.NewChatRepository,
			repository.NewMessageRepository,
			controller.NewChatController,
			controller.NewMessageController,
			handler.NewChatHandler,
			handler.NewMessageHandler,
			server.NewHTTPServer,
			server.NewServer,
		),
		fx.Invoke(registerServer),
		fx.Invoke(registerPostgres),
		fx.Invoke(registerGRPCClient),
	).Run()
}
