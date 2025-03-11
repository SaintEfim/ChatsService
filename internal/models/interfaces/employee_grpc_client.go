package interfaces

import (
	"context"

	"ChatsService/proto/employee"
)

type EmployeeGrpcClient interface {
	Initialize(ctx context.Context) error
	Close(ctx context.Context) error
	employee.GreeterEmployeesClient
}
