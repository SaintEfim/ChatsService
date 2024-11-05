package interfaces

import "github.com/gin-gonic/gin"

type Handler interface {
	ConfigureRoutes(r *gin.Engine)
	Get(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}
