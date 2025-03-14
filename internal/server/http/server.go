package http

import (
	"context"
	"errors"
	"net"
	"net/http"

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

func NewGRPCServer(chatGrpc interfaces.ChatGRPCServer) *grpc.Server {
	grpcServer := grpc.NewServer()
	chat.RegisterGreeterChatsServer(grpcServer, chatGrpc)
	return grpcServer
}

func NewServer(httpSrv *http.Server, grpcSrv *grpc.Server, cfg *config.Config,
	chatHandler interfaces.Handler[dto.Chat],
	messageHandler interfaces.Handler[dto.Message],
	logger *zap.Logger) interfaces.Server {
	return &Server{
		httpSrv:        httpSrv,
		grpcSrv:        grpcSrv,
		cfg:            cfg,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		logger:         logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		g := gin.Default()
		g.Use(middleware.LoggingMiddleware(s.logger))
		g.Use(middleware.AuthMiddleware(s.logger, s.cfg.AuthenticationConfiguration.AccessSecretKey))
		s.setGinMode(ctx)
		s.configureSwagger(ctx, g)
		s.configurationHandler(ctx, g)

		handler := CorsSettings(s.cfg).Handler(g)
		s.httpSrv.Handler = handler

		s.logger.Sugar().Infof("HTTP Server run on %s", s.cfg.HTTPServer.Port)
		if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("HTTP server error", zap.Error(err))
		}
	}()

	go func() {
		lis, err := net.Listen(s.cfg.GRPCServer.Type, s.cfg.GRPCServer.Addr)
		if err != nil {
			s.logger.Error("gRPC net.Listen error", zap.Error(err))
		}

		s.logger.Sugar().Infof("gRPC server run on %s", s.cfg.GRPCServer.Addr)
		if err := s.grpcSrv.Serve(lis); err != nil {
			s.logger.Error("gRPC server error", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, s.cfg.HTTPServer.Timeout)
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
