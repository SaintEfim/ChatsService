package handler

import (
	"fmt"
	"net/http"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"
	"ChatsService/internal/models/model"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

type ChatHandler struct {
	controller interfaces.Controller[model.ChatModel]
}

func NewChatHandler(controller interfaces.Controller[model.ChatModel]) interfaces.Handler[dto.ChatDto] {
	return &ChatHandler{controller: controller}
}

func (h *ChatHandler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/chats", h.Get)
	r.GET("/api/v1/chats/:id", h.GetOneById)
	r.POST("/api/v1/chats", h.Create)
	r.DELETE("/api/v1/chats/:id", h.Delete)
	r.PUT("/api/v1/chats/:id", h.Update)
}

// Get - godoc
// @Summary List chats
// @Description get chats
// @Tags chats
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{data=[]dto.ChatDto} "Successful response"
// @Failure 500 {object} dto.Response
// @Router /api/v1/chats [get]
func (h *ChatHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	chats, err := h.controller.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error retrieving chats: %v", err)})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: chats})
}

// GetOneById - godoc
// @Summary Get chat by ID
// @Description get chat by id
// @Tags chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} dto.Response{data=dto.ChatDto} "Successful response"
// @Failure 404 {object} dto.Response
// @Router /api/v1/chat/{id} [get]
func (h *ChatHandler) GetOneById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	chat, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: chat})
}

// Create - godoc
// @Summary Create a new chat
// @Description create chat
// @Tags chats
// @Accept json
// @Produce json
// @Param chat body dto.CreateChatDto true "Chat info"
// @Success 201 {object} dto.Response{data=dto.ChatDto} "Chat created successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/chats [post]
func (h *ChatHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		chatCreateDto dto.CreateChatDto
		chatModel     model.ChatModel
	)

	if err := c.ShouldBindJSON(&chatCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&chatCreateDto).To(&chatModel); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping chat: %v", err)})
		return
	}

	if err := h.controller.Create(ctx, &chatModel); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error creating chat: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: "Chat created successfully",
		Data:    chatModel.Id,
	})
}

// Delete - godoc
// @Summary Delete chat by ID
// @Description delete chat
// @Tags chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} dto.Response "Chat deleted successfully"
// @Failure 404 {object} dto.Response
// @Router /api/v1/chats/{id} [delete]
func (h *ChatHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.controller.Delete(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "Chat deleted successfully"})
}

// Update - godoc
// @Summary Update chat by ID
// @Description update chat`
// @Tags chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Param chat body dto.UpdateChatDto true "Chat info"
// @Success 200 {object} dto.Response{data=dto.ChatDto} "Chat updated successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/chats/{id} [put]
func (h *ChatHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var chatUpdateDto dto.UpdateChatDto
	var chatModel model.ChatModel

	if err := c.ShouldBindJSON(&chatUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&chatUpdateDto).To(&chatModel); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping chat: %v", err)})
		return
	}

	if err := h.controller.Update(ctx, id, &chatModel); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error updating chat: %v", err)})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "Chat updated successfully"})
}
