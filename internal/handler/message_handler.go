package handler

import (
	"net/http"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/models/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
)

type MessageHandler struct {
	controller interfaces.Controller[model.MessageModel]
}

func NewMessageHandler(controller interfaces.Controller[model.MessageModel]) interfaces.Handler[dto.MessageDto] {
	return &MessageHandler{controller: controller}
}

func (h *MessageHandler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/messages", h.Get)
	r.GET("/api/v1/messages/:id", h.GetOneById)
	r.POST("/api/v1/messages", h.Create)
	r.DELETE("/api/v1/messages/:id", h.Delete)
	r.PATCH("/api/v1/messages/:id", h.Update)
}

// Get @Summary List messages
// @Tags Messages
// @Accept json
// @Produce json
// @Success 200 {object} []dto.MessageDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/messages [get]
func (h *MessageHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	messages, err := h.controller.Get(ctx)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, messages)
}

// GetOneById @Summary Get message by ID
// @Tags Messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} dto.MessageDto
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Router /api/v1/messages/{id} [get]
func (h *MessageHandler) GetOneById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	message, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, message)
}

// Create @Summary Create a new message
// @Tags Messages
// @Accept json
// @Produce json
// @Param message body dto.CreateMessageDto true "Message info"
// @Success 201 {object} dto.CreateActionDto
// @Failure 400 {object} dto.ErrorDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/messages [post]
func (h *MessageHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	messageCreateDto := &dto.CreateMessageDto{}
	messageModel := &model.MessageModel{}

	if err := c.ShouldBindJSON(messageCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	automapper.MapLoose(messageCreateDto, messageModel)

	createItemId, err := h.controller.Create(ctx, messageModel)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.CreateActionDto{Id: createItemId})
}

// Delete @Summary Delete message by ID
// @Tags Messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Router /api/v1/messages/{id} [delete]
func (h *MessageHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := h.controller.Delete(ctx, id); err != nil {
		c.Error(err)
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
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/messages/{id} [patch]
func (h *MessageHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	messageUpdateDto := &dto.UpdateMessageDto{}
	messageModel := &model.MessageModel{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(messageUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	automapper.MapLoose(messageUpdateDto, messageModel)

	if err := h.controller.Update(ctx, id, messageModel); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
