package validator

import (
	"context"
	"errors"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	clientValidator "ChatsService/pkg/validator"

	"github.com/google/uuid"
)

type ChatValidator struct {
	chatRepo          interfaces.Repository[entity.Chat]
	employeeValidator *clientValidator.EmployeeValidator
}

func NewChatValidator(
	chatRepo interfaces.Repository[entity.Chat],
	employeeValidator *clientValidator.EmployeeValidator,
) *ChatValidator {
	return &ChatValidator{
		chatRepo:          chatRepo,
		employeeValidator: employeeValidator,
	}
}

func (v *ChatValidator) Validate(ctx context.Context, chat *dto.ChatCreate) error {
	if chat == nil {
		return errors.New("chat cannot be nil")
	}

	if len(chat.ParticipantIds) != 2 {
		return errors.New("private chat must have exactly 2 participants")
	}

	if hasDuplicates(chat.ParticipantIds) {
		return errors.New("duplicate participant IDs found")
	}

	if err := v.employeeValidator.ValidateEmployeesExist(ctx, chat.ParticipantIds); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	existingChats, err := v.chatRepo.Get(ctx)
	if err != nil {
		return fmt.Errorf("error checking existing chats: %w", err)
	}

	for _, ec := range existingChats {
		if uuidSlicesEqual(ec.ParticipantIds, chat.ParticipantIds) {
			return errors.New("chat with these participants already exists")
		}
	}

	return nil
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
