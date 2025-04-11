package validator

import (
	"context"
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

func TestChatValidator_Scenarios(t *testing.T) {
	type mockChatRepoResult struct {
		chats []*entity.Chat
		err   error
	}

	type mockEmployeeSearchRes struct {
		res *employee.SearchResponse
		err error
	}

	testCases := []struct {
		name            string
		chat            *dto.ChatCreate
		mockChatResult  mockChatRepoResult
		mockEmployeeRes mockEmployeeSearchRes
		expectedError   string
	}{
		{
			name: "Valid chat creation",
			chat: newTestChat(),
			mockChatResult: mockChatRepoResult{
				chats: nil,
			},
			mockEmployeeRes: mockEmployeeSearchRes{
				res: &employee.SearchResponse{
					Employees: []*employee.Employee{
						{Id: "00000000-0000-0000-0000-000000000001"},
						{Id: "00000000-0000-0000-0000-000000000002"},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockChatRepo := mocks.NewRepository[entity.Chat](t)
			mockEmployeeClient := mocks.NewEmployeeGrpcClient(t)

			mockChatRepo.On("Get", mock.Anything).Return(tc.mockChatResult.chats, tc.mockChatResult.err).Once()
			mockEmployeeClient.On("Search", mock.Anything, mock.Anything).
				Return(tc.mockEmployeeRes.res, tc.mockEmployeeRes.err).
				Once()

			validator := NewChatValidator(
				mockChatRepo,
				clientValidator.NewEmployeeValidator(mockEmployeeClient),
			)

			err := validator.Validate(context.Background(), tc.chat)

			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
