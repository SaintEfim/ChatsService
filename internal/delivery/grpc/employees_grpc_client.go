package grpc

import (
	"context"

	"ChatsService/config"
	"ChatsService/internal/models/interfaces"
	"ChatsService/proto/employee"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmployeeGrpcClient struct {
	client employee.GreeterEmployeesClient
	cfg    *config.Config
	conn   *grpc.ClientConn
}

func NewEmployeeGrpcClient(ctx context.Context, cfg *config.Config) interfaces.EmployeeGrpc {
	return &EmployeeGrpcClient{
		cfg: cfg,
	}
}

func (k *EmployeeGrpcClient) Initialize(ctx context.Context) error {
	conn, err := grpc.NewClient(k.cfg.GRPCClient.Services["employees"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	k.conn = conn
	k.client = employee.NewGreeterEmployeesClient(k.conn)

	return nil
}

func (k *EmployeeGrpcClient) Create(ctx context.Context, in *employee.EmployeeCreateRequest, opts ...grpc.CallOption) (*employee.EmployeeCreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k *EmployeeGrpcClient) Search(ctx context.Context, in *employee.SearchRequest, opts ...grpc.CallOption) (*employee.SearchResponse, error) {
	searchResponse, err := k.client.Search(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}

func (k *EmployeeGrpcClient) Close(ctx context.Context) error {
	if k.conn == nil {
		return nil
	}
	return k.conn.Close()
}
