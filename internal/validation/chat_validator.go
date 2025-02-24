package validation

import (
	"context"
	"fmt"

	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/employee"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ChatValidator struct {
	validate       *validator.Validate
	chatRepo       interfaces.Repository[entity.Chat]
	employeeClient interfaces.EmployeeGrpc
}

func NewChatValidator(
	chatRepo interfaces.Repository[entity.Chat],
	employeeClient interfaces.EmployeeGrpc,
) *ChatValidator {
	v := &ChatValidator{
		validate:       validator.New(),
		chatRepo:       chatRepo,
		employeeClient: employeeClient,
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

	ctx := context.Background()
	v.validatePrivateChat(sl, chat)
	v.validateGroupChat(sl, chat)
	v.validateEmployeesExist(ctx, sl, chat)
}

func (v *ChatValidator) validatePrivateChat(sl validator.StructLevel, chat dto.ChatCreate) {
	if chat.IsGroup {
		return
	}

	if chat.Name != "" {
		sl.ReportError(chat.Name, "Name", "", "name must be empty for private chats", "")
	}

	if len(chat.EmployeeIds) != 2 {
		sl.ReportError(
			chat.EmployeeIds,
			"EmployeeIDs",
			"",
			"private chat must have exactly 2 participants",
			"",
		)
		return
	}

	existing, err := v.chatRepo.Get(context.Background())
	if err != nil {
		sl.ReportError(
			chat.Name,
			"Name",
			"",
			fmt.Sprintf("error checking existing chats: %v", err),
			"",
		)
		return
	}

	for _, ec := range existing {
		if uuidSlicesEqual(ec.EmployeeIds, chat.EmployeeIds) {
			sl.ReportError(
				chat.EmployeeIds,
				"EmployeeIDs",
				"",
				"chat with these participants already exists",
				"",
			)
			break
		}
	}
}

func (v *ChatValidator) validateGroupChat(sl validator.StructLevel, chat dto.ChatCreate) {
	if !chat.IsGroup {
		return
	}

	if len(chat.EmployeeIds) < 1 {
		sl.ReportError(
			chat.EmployeeIds,
			"EmployeeIDs",
			"",
			"group chat must have at least 1 participant",
			"",
		)
	}
}

func (v *ChatValidator) validateEmployeesExist(
	ctx context.Context,
	sl validator.StructLevel,
	chat dto.ChatCreate,
) {
	employeeIdsStr := make([]string, len(chat.EmployeeIds))
	for i, id := range chat.EmployeeIds {
		employeeIdsStr[i] = id.String()
	}

	exist, err := v.employeeClient.Search(ctx, &employee.SearchRequest{Ids: employeeIdsStr})
	if err != nil {
		sl.ReportError(
			chat.EmployeeIds,
			"EmployeeIDs",
			"",
			fmt.Sprintf("employee check failed: %v", err),
			"",
		)
		return
	}

	if exist == nil {
		sl.ReportError(
			chat.EmployeeIds,
			"EmployeeIDs",
			"",
			"one or more employees do not exist",
			"",
		)
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
