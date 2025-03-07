package validator

import (
	"context"
	"fmt"

	"ChatsService/internal/models/dto"
	clientValidator "ChatsService/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MessageValidator struct {
	validate          *validator.Validate
	employeeValidator *clientValidator.EmployeeValidator
}

func NewMessageValidator(employeeValidator *clientValidator.EmployeeValidator) *MessageValidator {
	v := &MessageValidator{
		validate:          validator.New(),
		employeeValidator: employeeValidator,
	}
	v.registerCustomValidations()
	return v
}

func (v *MessageValidator) Validate(message *dto.MessageCreate) error {
	if message == nil {
		return fmt.Errorf("message cannot be nil")
	}
	return v.validate.Struct(message)
}

func (v *MessageValidator) registerCustomValidations() {
	v.validate.RegisterStructValidation(v.validateMessageStruct, dto.MessageCreate{})
}

func (v *MessageValidator) validateMessageStruct(sl validator.StructLevel) {
	message, ok := sl.Current().Interface().(dto.MessageCreate)
	if !ok {
		sl.ReportError(nil, "", "", "invalid message type", "")
		return
	}

	if message.SenderId == message.ReceiverId {
		sl.ReportError(message.ReceiverId, "ReceiverId", "ReceiverId", "notEqual", "SenderId and ReceiverId must be different")
		return
	}

	ctx := context.Background()
	employees := []uuid.UUID{message.SenderId, message.ReceiverId}
	v.employeeValidator.ValidateEmployeesExist(ctx, sl, employees)
}
