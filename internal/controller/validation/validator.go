package validation

import (
	"context"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

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

func ValidateChatCreateStruct(fl validator.StructLevel, chatRepository interfaces.Repository[entity.Chat]) {
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

func NewValidator(chatRepo interfaces.Repository[entity.Chat]) *validator.Validate {
	v := validator.New()
	v.RegisterStructValidation(func(fl validator.StructLevel) {
		ValidateChatCreateStruct(fl, chatRepo)
	}, &dto.ChatCreate{})
	return v
}
