package middleware

import (
	"net/http"
	"runtime/debug"

	"ChatsService/internal/exception"
	"ChatsService/internal/models/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ExceptionMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var (
			response   *dto.ErrorDto
			statusCode int
		)

		switch e := err.(type) {
		case *exception.NotFoundException:
			statusCode = e.Code
			response = &dto.ErrorDto{
				Status:      e.Code,
				Description: e.Description,
				StackTrace:  getStacktrace(),
			}
		default:
			statusCode = http.StatusInternalServerError
			response = &dto.ErrorDto{
				Status:      http.StatusInternalServerError,
				Description: "Internal server error",
			}

			logger.Error("Unhandled exception", zap.String("stackTrace", string(debug.Stack())))
		}

		c.JSON(statusCode, response)
		c.Abort()
	}
}

func getStacktrace() string {
	if gin.Mode() == "Development" {
		return string(debug.Stack())
	}

	return ""
}
