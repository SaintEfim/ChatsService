package server

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"ChatsService/config"
	"ChatsService/docs"
	"ChatsService/internal/middleware"
	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/chat"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	httpSrv        *http.Server
	grpcSrv        *grpc.Server
	cfg            *config.Config
	chatHandler    interfaces.Handler[dto.Chat]
	messageHandler interfaces.Handler[dto.Message]
	logger         *zap.Logger
}

func NewHTTPServer(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr: net.JoinHostPort(cfg.HTTPServer.Addr, cfg.HTTPServer.Port),
	}
}

func NewGRPCServer(chatGrpc interfaces.ChatGRPC) *grpc.Server {
	grpcServer := grpc.NewServer()
	chat.RegisterGreeterChatsServer(grpcServer, chatGrpc)
	return grpcServer
}

func NewServer(httpSrv *http.Server, grpcSrv *grpc.Server, cfg *config.Config,
	chatHandler interfaces.Handler[dto.Chat],
	messageHandler interfaces.Handler[dto.Message],
	logger *zap.Logger) interfaces.Server {
	return &Server{
		httpSrv: &http.Server{
			Addr: net.JoinHostPort(cfg.HTTPServer.Addr, cfg.HTTPServer.Port),
		},
		grpcSrv:        grpcSrv,
		cfg:            cfg,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		logger:         logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		g := gin.Default()
		g.Use(middleware.LoggingMiddleware(s.logger))
		s.setGinMode(ctx)
		s.configureSwagger(ctx, g)
		s.configurationHandler(ctx, g)

		s.httpSrv.Handler = g

		s.logger.Sugar().Infof("HTTP Server run on %s", s.httpSrv.Addr)
		if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		lis, err := net.Listen(s.cfg.GRPCServer.Type, s.cfg.GRPCServer.Addr)
		if err != nil {
			errChan <- err
			return
		}
		s.logger.Sugar().Infof("gRPC server run %s", s.cfg.GRPCServer.Addr)
		if err := s.grpcSrv.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.httpSrv.Shutdown(shutdownCtx); err != nil {
			s.logger.Error("Ошибка при завершении HTTP сервера", zap.Error(err))
		}
		s.grpcSrv.GracefulStop()
	}

	wg.Wait()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.httpSrv.Shutdown(shutdownCtx); err != nil {
		return err
	}
	s.grpcSrv.GracefulStop()
	return nil
}

func (s *Server) configureSwagger(ctx context.Context, router *gin.Engine) {
	docs.SwaggerInfo.Title = "Chats Service API"
	docs.SwaggerInfo.Description = "This is a sample server Chats server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) setGinMode(ctx context.Context) {
	switch s.cfg.EnvironmentVariables.Environment {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		s.logger.Sugar().Infof("Unknown environment: %s, defaulting to 'development'", s.cfg.EnvironmentVariables.Environment)
		gin.SetMode(gin.DebugMode)
	}
}

func (s *Server) configurationHandler(ctx context.Context, g *gin.Engine) {
	s.chatHandler.ConfigureRoutes(g)
	s.messageHandler.ConfigureRoutes(g)
}
