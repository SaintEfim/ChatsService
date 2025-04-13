package validator

import (
	"context"
	"errors"
	"testing"

	"ChatsService/internal/mocks"
	"ChatsService/internal/models/dto"
	"ChatsService/internal/models/entity"
	clientValidator "ChatsService/pkg/validator"
	"ChatsService/proto/employee"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newTestChat(modifiers ...func(*dto.ChatCreate)) *dto.ChatCreate {
	base := &dto.ChatCreate{
		ParticipantIds: []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		},
	}

	for _, modify := range modifiers {
		modify(base)
	}
	return base
}

func TestChatValidator_InputValidation(t *testing.T) {
	tests := []struct {
		name          string
		chat          *dto.ChatCreate
		expectedError error
	}{
		{
			name: "More than two participants",
			chat: newTestChat(func(create *dto.ChatCreate) {
				create.ParticipantIds = []uuid.UUID{
					uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				}
			}),
			expectedError: errors.New("private chat must have exactly 2 participants"),
		},
		{
			name: "Chat has duplicate participants",
			chat: newTestChat(func(create *dto.ChatCreate) {
				create.ParticipantIds = []uuid.UUID{
					uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				}
			}),
			expectedError: errors.New("duplicate participant IDs found"),
		},
		{
			name:          "Nil chat",
			chat:          nil,
			expectedError: errors.New("chat cannot be nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewChatValidator(nil, nil)
			err := validator.Validate(context.Background(), tt.chat)
			assertError(t, err, tt.expectedError)
		})
	}
}

func TestChatValidator_EmployeeErrors(t *testing.T) {
	tests := []struct {
		name          string
		mockResponse  *employee.SearchResponse
		mockError     error
		expectedError error
	}{
		{
			name: "Participant not found",
			mockResponse: &employee.SearchResponse{
				Employees: []*employee.Employee{
					{Id: "00000000-0000-0000-0000-000000000001"},
				},
			},
			expectedError: errors.New("employee validation failed: one or more employees do not exist"),
		},
		{
			name:          "gRPC error",
			mockError:     errors.New("connection error"),
			expectedError: errors.New("employee validation failed: employee check failed: connection error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChatRepo := mocks.NewRepository[entity.Chat](t)
			mockEmployeeClient := mocks.NewEmployeeGrpcClient(t)

			// Настройка моков
			mockEmployeeClient.On("Search", mock.Anything, expectedSearchRequest()).
				Return(tt.mockResponse, tt.mockError).
				Once()

			validator := NewChatValidator(
				mockChatRepo,
				clientValidator.NewEmployeeValidator(mockEmployeeClient),
			)

			err := validator.Validate(context.Background(), newTestChat())
			assertError(t, err, tt.expectedError)

			mockEmployeeClient.AssertExpectations(t)
		})
	}
}

func TestChatValidator_WithValidEmployee(t *testing.T) {
	tests := []struct {
		name           string
		mockChatResult []*entity.Chat
		mockChatError  error
		expectedError  error
	}{
		{
			name:           "Valid chat creation",
			mockChatResult: []*entity.Chat{},
		},
		{
			name:          "Chat repository error",
			mockChatError: errors.New("database error"),
			expectedError: errors.New("error checking existing chats: database error"),
		},
		{
			name: "Chat already exists",
			mockChatResult: []*entity.Chat{
				{
					ParticipantIds: []uuid.UUID{
						uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					},
				},
			},
			expectedError: errors.New("chat with these participants already exists"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockChatRepo := mocks.NewRepository[entity.Chat](t)
			mockEmployeeClient := mocks.NewEmployeeGrpcClient(t)

			// Настройка моков
			mockChatRepo.On("Get", mock.Anything).
				Return(tt.mockChatResult, tt.mockChatError).
				Once()

			mockEmployeeClient.On("Search", mock.Anything, expectedSearchRequest()).
				Return(validEmployeesResponse(), nil).
				Once()

			validator := NewChatValidator(
				mockChatRepo,
				clientValidator.NewEmployeeValidator(mockEmployeeClient),
			)

			err := validator.Validate(context.Background(), newTestChat())
			assertError(t, err, tt.expectedError)

			mockChatRepo.AssertExpectations(t)
			mockEmployeeClient.AssertExpectations(t)
		})
	}
}

func expectedSearchRequest() *employee.SearchRequest {
	return &employee.SearchRequest{
		Ids: []string{
			"00000000-0000-0000-0000-000000000001",
			"00000000-0000-0000-0000-000000000002",
		},
	}
}

func validEmployeesResponse() *employee.SearchResponse {
	return &employee.SearchResponse{
		Employees: []*employee.Employee{
			{Id: "00000000-0000-0000-0000-000000000001"},
			{Id: "00000000-0000-0000-0000-000000000002"},
		},
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()
	if expected != nil {
		require.Error(t, actual)
		assert.Contains(t, actual.Error(), expected.Error())
	} else {
		assert.NoError(t, actual)
	}
}
