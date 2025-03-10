package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"ChatsService/config"
	"ChatsService/docs"
	"ChatsService/internal/middleware"
	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
)

type Server struct {
	srv            *http.Server
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

func NewServer(srv *http.Server, cfg *config.Config,
	chatHandler interfaces.Handler[dto.Chat],
	messageHandler interfaces.Handler[dto.Message],
	logger *zap.Logger) interfaces.Server {
	return &Server{
		srv:            srv,
		cfg:            cfg,
		chatHandler:    chatHandler,
		messageHandler: messageHandler,
		logger:         logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	var err error

	go func() {
		g := gin.Default()

		g.Use(middleware.LoggingMiddleware(s.logger))

		s.setGinMode(ctx)
		s.configureSwagger(ctx, g)
		s.configurationHandler(ctx, g)

		handler := CorsSettings(s.cfg).Handler(g)

		s.srv.Handler = handler

		s.logger.Sugar().Infof("Listening and serving HTTP on %s\n", s.srv.Addr)

		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(err.Error())
		}
	}()

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	s.srv.RegisterOnShutdown(cancel)

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

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
