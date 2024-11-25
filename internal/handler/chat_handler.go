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

// Get @Summary List chats
// @Tags Chats
// @Accept json
// @Produce json
// @Success 200 {object} []dto.ChatDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/chats [get]
func (h *ChatHandler) Get(c *gin.Context) {
	chatDtos := make([]dto.ChatDto, 0)
	ctx := c.Request.Context()

	chats, err := h.controller.Get(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	automapper.MapLoose(chats, &chatDtos)

	c.JSON(http.StatusOK, chatDtos)
}

// GetOneById @Summary Get chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} dto.ChatDto
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Router /api/v1/chats/{id} [get]
func (h *ChatHandler) GetOneById(c *gin.Context) {
	chatDto := &dto.ChatDto{}
	ctx := c.Request.Context()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	chat, err := h.controller.GetOneById(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}

	automapper.MapLoose(chat, chatDto)
	c.JSON(http.StatusOK, chatDto)
}

// Create @Summary Create a new chat
// @Tags Chats
// @Accept json
// @Produce json
// @Param chat body dto.CreateChatDto true "Chat info"
// @Success 201 {object} dto.CreateActionDto
// @Failure 400 {object} dto.ErrorDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/chats [post]
func (h *ChatHandler) Create(c *gin.Context) {
	chatCreateDto := &dto.CreateChatDto{}
	chatModel := &model.ChatModel{}
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(chatCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	automapper.MapLoose(chatCreateDto, chatModel)

	createItemId, err := h.controller.Create(ctx, chatModel)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.CreateActionDto{Id: createItemId})
}

// Delete @Summary Delete chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Router /api/v1/chats/{id} [delete]
func (h *ChatHandler) Delete(c *gin.Context) {
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

// Update @Summary Update chat by ID
// @Tags Chats
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Param chat body dto.UpdateChatDto true "Chat info"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorDto
// @Failure 404 {object} dto.ErrorDto
// @Failure 500 {object} dto.ErrorDto
// @Router /api/v1/chats/{id} [put]
func (h *ChatHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	chatUpdateDto := &dto.UpdateChatDto{}
	chatModel := &model.ChatModel{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&chatUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	automapper.MapLoose(chatUpdateDto, chatModel)

	if err := h.controller.Update(ctx, id, chatModel); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
