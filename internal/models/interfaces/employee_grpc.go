package interfaces

import (
	"context"

	"ChatsService/proto/employee"
)

type EmployeeGrpc interface {
	Initialize(ctx context.Context) error
	Close(ctx context.Context) error
	employee.GreeterEmployeesClient
}
