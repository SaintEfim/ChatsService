package middleware

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		logFields := []zap.Field{
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("header", fmt.Sprintf("%v", c.Request.Header)),
			zap.String("query_parameters", c.Request.URL.Query().Encode()),
			zap.String("size_request", fmt.Sprintf("%d", c.Writer.Size())),
		}

		requestBody, err := readRequestBody(c.Request.Body)

		if err != nil {
			logger.Error("Failed to read request body", zap.Error(err))
		} else {
			logFields = append(logFields, zap.String("request_body", requestBody))
		}

		logger.Info("Incoming request",
			logFields...,
		)

		c.Next()

		if len(c.Errors) > 0 {
			logFields = append(logFields, zap.String("error", c.Errors.String()))
			logger.Error("Request completed with errors", logFields...)
		} else {
			logFields = append(logFields,
				zap.Int("status", c.Writer.Status()),
				zap.Duration("duration", time.Since(startTime)),
			)

			logger.Info("Request completed", logFields...)
		}
	}
}

func readRequestBody(body io.ReadCloser) (string, error) {
	if body == nil {
		return "", errors.New("body is not empty")
	}
	defer body.Close()
	buf, err := io.ReadAll(body)

	if err != nil {
		return "", err
	}
	return string(buf), nil
}
