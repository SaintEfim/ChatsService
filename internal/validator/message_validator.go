package validator

import (
	"context"
	"errors"
	"fmt"

	"ChatsService/internal/models/dto"
	clientValidator "ChatsService/pkg/validator"

	"github.com/google/uuid"
)

type MessageValidator struct {
	employeeValidator *clientValidator.EmployeeValidator
}

func NewMessageValidator(employeeValidator *clientValidator.EmployeeValidator) *MessageValidator {
	return &MessageValidator{
		employeeValidator: employeeValidator,
	}
}

func (v *MessageValidator) Validate(message *dto.MessageCreate) error {
	if message == nil {
		return errors.New("message cannot be nil")
	}

	if message.Text == "" {
		return errors.New("text is required")
	}

	if message.SenderId == message.ReceiverId {
		return errors.New("sender and receiver must be different")
	}

	ctx := context.Background()
	employees := []uuid.UUID{message.SenderId, message.ReceiverId}
	if err := v.employeeValidator.ValidateEmployeesExist(ctx, employees); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	return nil
}
