package validator

import (
	"context"
	"fmt"

	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/employee"

	"github.com/go-playground/validator/v10"
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
	sl validator.StructLevel,
	employeeIds []uuid.UUID,
) {
	employeeIdsStr := make([]string, len(employeeIds))
	for i, id := range employeeIds {
		employeeIdsStr[i] = id.String()
	}

	exist, err := v.client.Search(ctx, &employee.SearchRequest{Ids: employeeIdsStr})
	if err != nil {
		sl.ReportError(
			"employeeId",
			"EmployeeIds",
			"",
			fmt.Sprintf("employee check failed: %v", err),
			"",
		)
		return
	}

	if exist == nil {
		sl.ReportError(
			"employeeId",
			"EmployeeIds",
			"",
			"not found employees",
			"",
		)
	}
}
