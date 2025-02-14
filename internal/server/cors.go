package server

import (
	"net/http"

	"github.com/rs/cors"

	"ChatsService/config"
)

func CorsSettings(cfg *config.Config) *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodOptions, http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedOrigins: cfg.Cors.AllowedOrigins,
		AllowedHeaders: []string{
			"Content-Type",
		},
		Debug: true,
	})

	return c
}
