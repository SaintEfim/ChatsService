package middleware

import (
	"net/http"
	"strings"

	"ChatsService/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func AuthMiddleware(logger *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/swagger") {
			c.Next()
			return
		}
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.AuthenticationConfiguration.AccessSecretKey), nil
		})

		if err != nil || !token.Valid {
			logger.Error("Invalid token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user", token.Claims)

		c.Next()
	}
}
