package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
)

type MessageHandler struct {
	controller interfaces.MessageController
}

func NewMessageHandler(controller interfaces.MessageController) interfaces.Handler[dto.Message] {
	return &MessageHandler{controller: controller}
}

func (h *MessageHandler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/messages", h.Get)
	r.GET("/api/v1/messages/:id", h.GetOneById)
	r.POST("/api/v1/messages", h.Create)
	r.DELETE("/api/v1/messages/:id", h.Delete)
	r.PUT("/api/v1/messages/:id", h.Update)
}

// Get @Summary List messages
// @Tags Messages
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Message
// @Failure 500 {object} dto.Error
// @Router /api/v1/messages [get]
func (h *MessageHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	messages, err := h.controller.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// GetOneById @Summary Get message by ID
// @Tags Messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Router /api/v1/messages/{id} [get]
func (h *MessageHandler) GetOneById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	message, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, message)
}

// Create @Summary Create a new message
// @Tags Messages
// @Accept json
// @Produce json
// @Param message body dto.CreateMessageDto true "Message info"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /api/v1/messages [post]
func (h *MessageHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	messageCreate := &dto.MessageCreate{}

	if err := c.ShouldBindJSON(messageCreate); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	createItem, err := h.controller.Create(ctx, messageCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateAction{Id: createItem.Id})
}

// Delete @Summary Delete message by ID
// @Tags Messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Router /api/v1/messages/{id} [delete]
func (h *MessageHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := h.controller.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// Update @Summary Update message by ID
// @Tags Messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Param message body dto.UpdateMessageDto true "Message info"
// @Success 204 "No Content"
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /api/v1/messages/{id} [put]
func (h *MessageHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	messageUpdate := &dto.MessageUpdate{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(messageUpdate); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := h.controller.Update(ctx, id, messageUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
