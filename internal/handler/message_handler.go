package handler

import (
	"fmt"
	"net/http"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

type MessageHandler struct {
	controller interfaces.Controller[entity.MessageEntity]
}

func NewMessageHandler(controller interfaces.Controller[entity.MessageEntity]) interfaces.Handler {
	return &MessageHandler{controller: controller}
}

func (h *MessageHandler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/messages", h.Get)
	r.GET("/api/v1/messages/:id", h.GetOneById)
	r.POST("/api/v1/messages", h.Create)
	r.DELETE("/api/v1/messages/:id", h.Delete)
	r.PUT("/api/v1/messages/:id", h.Update)
}

// Get - godoc
// @Summary List messages
// @Description get messages
// @Tags messages
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{data=[]dto.MessageDto} "Successful response"
// @Failure 500 {object} dto.Response
// @Router /api/v1/messages [get]
func (h *MessageHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	messages, err := h.controller.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error retrieving messages: %v", err)})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: messages})
}

// GetOneById - godoc
// @Summary Get message by ID
// @Description get message by id
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} dto.Response{data=dto.MessageDto} "Successful response"
// @Failure 404 {object} dto.Response
// @Router /api/v1/message/{id} [get]
func (h *MessageHandler) GetOneById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	message, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "Message not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: message})
}

// Create - godoc
// @Summary Create a new message
// @Description create message
// @Tags messages
// @Accept json
// @Produce json
// @Param message body dto.CreateMessageDto true "Message info"
// @Success 201 {object} dto.Response{data=dto.MessageDto} "Message created successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/messages [post]
func (h *MessageHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		messageCreateDto dto.CreateMessageDto
		messageEntity    entity.MessageEntity
	)

	if err := c.ShouldBindJSON(&messageCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&messageCreateDto).To(&messageEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping message: %v", err)})
		return
	}

	if err := h.controller.Create(ctx, &messageEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error creating message: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: "Message created successfully",
		Data:    messageEntity.Id,
	})
}

// Delete - godoc
// @Summary Delete message by ID
// @Description delete message
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} dto.Response "Message deleted successfully"
// @Failure 404 {object} dto.Response
// @Router /api/v1/messages/{id} [delete]
func (h *MessageHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.controller.Delete(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "Message not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "Message deleted successfully"})
}

// Update - godoc
// @Summary Update message by ID
// @Description update message`
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Param message body dto.UpdateMessageDto true "Message info"
// @Success 200 {object} dto.Response{data=dto.MessageDto} "Message updated successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/messages/{id} [put]
func (h *MessageHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var messageUpdateDto dto.UpdateMessageDto
	var messageEntity entity.MessageEntity

	if err := c.ShouldBindJSON(&messageUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&messageUpdateDto).To(&messageEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping message: %v", err)})
		return
	}

	if err := h.controller.Update(ctx, id, &messageEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error updating message: %v", err)})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "Message updated successfully"})
}