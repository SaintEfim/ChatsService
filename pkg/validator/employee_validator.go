package validator

import (
	"context"
	"errors"
	"fmt"

	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/employee"

	"github.com/google/uuid"
)

type EmployeeValidator struct {
	client interfaces.EmployeeGrpcClient
}

func NewEmployeeValidator(client interfaces.EmployeeGrpcClient) *EmployeeValidator {
	return &EmployeeValidator{
		client: client,
	}
}

func (v *EmployeeValidator) ValidateEmployeesExist(
	ctx context.Context,
	employeeIds []uuid.UUID,
) error {
	employeeIdsStr := make([]string, len(employeeIds))
	for i, id := range employeeIds {
		employeeIdsStr[i] = id.String()
	}

	exist, err := v.client.Search(ctx, &employee.SearchRequest{Ids: employeeIdsStr})
	if err != nil {
		return fmt.Errorf("employee check failed: %w", err)
	}

	if exist == nil || len(exist.Employees) != len(employeeIds) {
		return errors.New("one or more employees do not exist")
	}

	return nil
}
