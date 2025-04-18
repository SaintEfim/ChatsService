package handler

import (
	"net/http"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatHandler struct {
	controller interfaces.ChatController
}

func NewChatHandler(controller interfaces.ChatController) interfaces.Handler[dto.Chat] {
	return &ChatHandler{controller: controller}
}

func (h *ChatHandler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/chats", h.Get)
	r.GET("/api/v1/chats/user/:id", h.GetChatsByUserId)
	r.GET("/api/v1/chats/:id", h.GetOneById)
	r.POST("/api/v1/chats", h.Create)
	r.DELETE("/api/v1/chats/:id", h.Delete)
}

// Get @Summary List chats
// @Tags Chats
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Chat
// @Failure 500 {object} dto.Error
// @Security BearerAuth
// @Router /api/v1/chats [get]
func (h *ChatHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	chats, err := h.controller.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chats)
}

// GetChatsByUserId @Summary Get chats by User ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} []dto.Chat
// @Failure 500 {object} dto.Error
// @Security BearerAuth
// @Router /api/v1/chats/user/{id} [get]
func (h *ChatHandler) GetChatsByUserId(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	chats, err := h.controller.GetChatsByUserId(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chats)
}

// GetOneById @Summary Get chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} dto.ChatDetail
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Security BearerAuth
// @Router /api/v1/chats/{id} [get]
func (h *ChatHandler) GetOneById(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	chat, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chat)
}

// Create @Summary Create a new chat
// @Tags Chats
// @Accept json
// @Produce json
// @Param chat body dto.ChatCreate true "Chat info"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Security BearerAuth
// @Router /api/v1/chats [post]
func (h *ChatHandler) Create(c *gin.Context) {
	chatCreate := &dto.ChatCreate{}
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(chatCreate); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	createItem, err := h.controller.Create(ctx, chatCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &dto.CreateAction{Id: createItem.Id})
}

func (h *ChatHandler) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Delete @Summary Delete chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Security BearerAuth
// @Router /api/v1/chats/{id} [delete]
func (h *ChatHandler) Delete(c *gin.Context) {
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
