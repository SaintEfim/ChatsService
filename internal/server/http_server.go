package server

import (
	"fmt"

	"ChatsService/config"

	"net/http"
)

func NewHTTPServer(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.HTTPServer.Addr, cfg.HTTPServer.Port),
	}
}
