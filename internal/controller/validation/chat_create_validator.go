package validation

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ChatCreateValidator struct {
	validate *validator.Validate
	chatRepo interfaces.Repository[entity.Chat]
}

func NewChatCreateValidator(chatRepo interfaces.Repository[entity.Chat]) *ChatCreateValidator {
	validator := &ChatCreateValidator{
		validate: validator.New(),
		chatRepo: chatRepo,
	}
	validator.registerCustomValidations()
	return validator
}

func (v *ChatCreateValidator) registerCustomValidations() {
	v.validate.RegisterStructValidation(func(fl validator.StructLevel) {
		v.validateChatCreateStruct(fl, v.chatRepo)
	}, &dto.ChatCreate{})
}

func (v *ChatCreateValidator) ValidateStruct(chat *dto.ChatCreate) error {
	return v.validate.Struct(chat)
}

func reportError(fl validator.StructLevel, field interface{}, fieldName, message string) {
	fl.ReportError(field, fieldName, "", message, "")
}

func areUUIDSlicesEqual(a, b []uuid.UUID) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (v *ChatCreateValidator) validateChatCreateStruct(fl validator.StructLevel, chatRepository interfaces.Repository[entity.Chat]) {
	ctx := context.Background()
	chat := fl.Current().Interface().(dto.ChatCreate)

	if !chat.IsGroup {
		if chat.Name != "" {
			reportError(fl, chat.Name, "Name", "name must be empty for private chats")
		}

		if len(chat.EmployeeIds) != 2 {
			reportError(fl, chat.EmployeeIds, "EmployeeIds", "private chat must have exactly 2 participants")
		}
	} else {
		if len(chat.EmployeeIds) < 1 {
			reportError(fl, chat.EmployeeIds, "EmployeeIds", "group chat must have at least 1 participant")
		}
	}

	existingChats, err := chatRepository.Get(ctx)
	if err != nil {
		reportError(fl, chat.Name, "Name", "error checking existing chats: "+err.Error())
		return
	}

	for _, existingChat := range existingChats {
		if areUUIDSlicesEqual(existingChat.EmployeeIds, chat.EmployeeIds) {
			reportError(fl, chat.EmployeeIds, "EmployeeIds", "chat with these participants already exists")
			return
		}
	}
}
