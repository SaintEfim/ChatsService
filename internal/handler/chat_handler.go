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
	r.GET("/api/v1/chats/user/:id/interlocutor/:interlocutorId", h.PrivateChatExists)
	r.GET("/api/v1/chats/:id", h.GetOneById)
	r.POST("/api/v1/chats", h.Create)
	r.DELETE("/api/v1/chats/:id", h.Delete)
	r.PUT("/api/v1/chats/:id", h.Update)
}

// Get @Summary List chats
// @Tags Chats
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Chat
// @Failure 500 {object} dto.Error
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

// PrivateChatExists checks if a chat exists between two users
// @Summary Check if a chat exists
// @Tags Chats
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param colleague_id path string true "Colleague ID"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /api/v1/chats/user/{user_id}/colleague/{colleague_id} [get]
func (h *ChatHandler) PrivateChatExists(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	interlocutorId, err := uuid.Parse(c.Param("interlocutorId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	exists, err := h.controller.PrivateChatExists(ctx, userId, interlocutorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

// GetOneById @Summary Get chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} dto.ChatDetail
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
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

// Delete @Summary Delete chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
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

// Update @Summary Update chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Param chat body dto.ChatUpdate true "Chat info"
// @Success 204 "No Content"
// @Failure 400 {object} dto.Error
// @Failure 404 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /api/v1/chats/{id} [put]
func (h *ChatHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	chatUpdate := &dto.ChatUpdate{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&chatUpdate); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := h.controller.Update(ctx, id, chatUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:      http.StatusInternalServerError,
			Description: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
