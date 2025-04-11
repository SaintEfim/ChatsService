package validator

import (
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

func NewChatCreate(modifiers ...func(*dto.ChatCreate)) *dto.ChatCreate {
	base := &dto.ChatCreate{
		Name: "",
		ParticipantIds: []uuid.UUID{
			uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		},
	}

	msg := base
	for _, modify := range modifiers {
		modify(msg)
	}

	return msg
}

func TestChatValidator(t *testing.T) {
	type mockChatRepositoryResult struct {
		response []*entity.Chat
		err      error
	}

	type mockEmployeeSearchResult struct {
		response *employee.SearchResponse
		err      error
	}

	tests := []struct {
		name                     string
		chat                     *dto.ChatCreate
		mockChatRepositoryResult mockChatRepositoryResult
		employeeSearchResult     mockEmployeeSearchResult
		expectErrContains        string
	}{
		{
			name: "Valid chat",
			chat: NewChatCreate(),
			mockChatRepositoryResult: mockChatRepositoryResult{
				response: nil,
			},
			employeeSearchResult: mockEmployeeSearchResult{
				response: &employee.SearchResponse{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockChatRepository := mocks.NewRepository[entity.Chat](t)
			mockEmployeeGrpcClient := mocks.NewEmployeeGrpcClient(t)

			mockChatRepository.On("Get", mock.Anything).Return(tc.mockChatRepositoryResult.response, tc.mockChatRepositoryResult.err).Once()

			mockEmployeeGrpcClient.On("Search", mock.Anything, mock.Anything).
				Return(tc.employeeSearchResult.response, tc.employeeSearchResult.err).
				Once()

			validator := NewChatValidator(
				mockChatRepository,
				clientValidator.NewEmployeeValidator(mockEmployeeGrpcClient),
			)

			err := validator.Validate(tc.chat)

			if tc.expectErrContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
