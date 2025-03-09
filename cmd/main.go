// @title ChatsService API
// @version 1.0
// @host localhost:1006
// @BasePath /
// @schemes http https
package main

import (
	"context"

	"ChatsService/config"
	"ChatsService/internal/controller"
	"ChatsService/internal/delivery/grpc"
	"ChatsService/internal/handler"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/postgres"
	"ChatsService/internal/repository"
	"ChatsService/internal/server"
	"ChatsService/internal/validator"
	"ChatsService/pkg/logger"
	clientValidator "ChatsService/pkg/validator"

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
		OnStart: srv.Run,
		OnStop:  srv.Stop,
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
			validator.NewChatValidator,
			validator.NewMessageValidator,
			repository.NewChatRepository,
			repository.NewMessageRepository,
			controller.NewChatController,
			controller.NewMessageController,
			handler.NewChatHandler,
			handler.NewMessageHandler,
			server.NewHTTPServer,
			server.NewServer,
			clientValidator.NewEmployeeValidator,
		),
		fx.Invoke(registerServer),
		fx.Invoke(registerPostgres),
		fx.Invoke(registerGRPCClient),
	).Run()
}
