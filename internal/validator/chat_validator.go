package validator

import (
	"context"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	clientValidator "ChatsService/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ChatValidator struct {
	validate          *validator.Validate
	chatRepo          interfaces.Repository[entity.Chat]
	employeeValidator *clientValidator.EmployeeValidator
}

func NewChatValidator(chatRepo interfaces.Repository[entity.Chat], employeeValidator *clientValidator.EmployeeValidator) *ChatValidator {
	v := &ChatValidator{
		validate:          validator.New(),
		chatRepo:          chatRepo,
		employeeValidator: employeeValidator,
	}
	v.registerCustomValidations()
	return v
}

func (v *ChatValidator) Validate(chat *dto.ChatCreate) error {
	if chat == nil {
		return fmt.Errorf("chat cannot be nil")
	}
	return v.validate.Struct(chat)
}

func (v *ChatValidator) registerCustomValidations() {
	v.validate.RegisterStructValidation(v.validateChatStruct, dto.ChatCreate{})
}

func (v *ChatValidator) validateChatStruct(sl validator.StructLevel) {
	chat, ok := sl.Current().Interface().(dto.ChatCreate)
	if !ok {
		sl.ReportError(nil, "", "", "invalid chat type", "")
		return
	}

	if hasDuplicates(chat.ParticipantIds) {
		sl.ReportError(chat.ParticipantIds, "ParticipantIds", "ParticipantIds", "duplicate participant IDs found", "")
		return
	}

	ctx := context.Background()
	v.validatePrivateChat(sl, chat)
	v.validateGroupChat(sl, chat)
	v.employeeValidator.ValidateEmployeesExist(ctx, sl, chat.ParticipantIds)
}

func (v *ChatValidator) validatePrivateChat(sl validator.StructLevel, chat dto.ChatCreate) {
	if chat.Name != "" {
		sl.ReportError(chat.Name, "Name", "", "name must be empty for private chats", "")
	}
	if len(chat.ParticipantIds) != 2 {
		sl.ReportError(chat.ParticipantIds, "ParticipantIds", "", "private chat must have exactly 2 participants", "")
		return
	}

	existing, err := v.chatRepo.Get(context.Background())
	if err != nil {
		sl.ReportError(chat.Name, "Name", "", fmt.Sprintf("error checking existing chats: %v", err), "")
		return
	}
	for _, ec := range existing {
		if uuidSlicesEqual(ec.ParticipantIds, chat.ParticipantIds) {
			sl.ReportError(chat.ParticipantIds, "ParticipantIds", "", "chat with these participants already exists", "")
			break
		}
	}
}

func (v *ChatValidator) validateGroupChat(sl validator.StructLevel, chat dto.ChatCreate) {
	if len(chat.ParticipantIds) < 1 {
		sl.ReportError(chat.ParticipantIds, "ParticipantIds", "", "group chat must have at least 1 participant", "")
	}
}

func uuidSlicesEqual(a, b []uuid.UUID) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[uuid.UUID]struct{}, len(a))
	for _, id := range a {
		seen[id] = struct{}{}
	}
	for _, id := range b {
		if _, exists := seen[id]; !exists {
			return false
		}
	}
	return true
}

func hasDuplicates(ids []uuid.UUID) bool {
	seen := make(map[uuid.UUID]struct{}, len(ids))
	for _, id := range ids {
		if _, exists := seen[id]; exists {
			return true
		}
		seen[id] = struct{}{}
	}
	return false
}
