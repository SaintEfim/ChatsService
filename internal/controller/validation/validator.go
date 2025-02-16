package validation

import (
	"github.com/go-playground/validator/v10"

	"ChatsService/internal/models/dto"
)

func reportError(fl validator.StructLevel, field interface{}, fieldName, message string) {
	fl.ReportError(field, fieldName, "", message, "")
}

func ValidateChatCreateStruct(fl validator.StructLevel) {
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
			reportError(fl, chat.EmployeeIds, "EmployeeIds", "group chat must have at least 1 participants")
		}
	}
}

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterStructValidation(ValidateChatCreateStruct, &dto.ChatCreate{})
	return v
}
