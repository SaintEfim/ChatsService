package validator

import (
	"errors"
	"testing"

	"ChatsService/internal/mocks"
	"ChatsService/internal/models/dto"
	clientValidator "ChatsService/pkg/validator"
	"ChatsService/proto/employee"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func NewMessageCreate(modifiers ...func(*dto.MessageCreate)) *dto.MessageCreate {
	base := &dto.MessageCreate{
		ChatId:     uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		SenderId:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		ReceiverId: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
		Text:       "Hello",
	}

	msg := base
	for _, modify := range modifiers {
		modify(msg)
	}

	return msg
}

func TestMessageValidatorSenderIdEqualReceiverId(t *testing.T) {
	mockClient := mocks.NewEmployeeGrpcClient(t)
	validator := NewMessageValidator(clientValidator.NewEmployeeValidator(mockClient))
	expectErrContains := "failed on the 'notEqual' tag"

	err := validator.Validate(NewMessageCreate(func(m *dto.MessageCreate) {
		m.ReceiverId = m.SenderId
	}))

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectErrContains)

	mockClient.AssertNotCalled(t, "Search", mock.Anything, mock.Anything)
}

func TestMessageValidator(t *testing.T) {
	type mockSearchResult struct {
		response *employee.SearchResponse
		err      error
	}

	testCases := []struct {
		name              string
		message           *dto.MessageCreate
		searchResult      mockSearchResult
		expectErrContains string
	}{
		{
			name:    "Valid message",
			message: NewMessageCreate(),
			searchResult: mockSearchResult{
				response: &employee.SearchResponse{},
			},
		},
		{
			name:    "Employee search returns error",
			message: NewMessageCreate(),
			searchResult: mockSearchResult{
				err: errors.New("employee check failed"),
			},
			expectErrContains: "employee check failed",
		},
		{
			name:    "Employee not found",
			message: NewMessageCreate(),
			searchResult: mockSearchResult{
				response: nil,
			},
			expectErrContains: "not found employees",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockEmployeeGrpcClient := mocks.NewEmployeeGrpcClient(t)

			mockEmployeeGrpcClient.On("Search", mock.Anything, mock.Anything).
				Return(tc.searchResult.response, tc.searchResult.err).
				Once()

			validator := NewMessageValidator(
				clientValidator.NewEmployeeValidator(mockEmployeeGrpcClient),
			)

			err := validator.Validate(tc.message)

			if tc.expectErrContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErrContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
