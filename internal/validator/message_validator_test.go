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

	err := validator.Validate(NewMessageCreate(func(m *dto.MessageCreate) {
		m.ReceiverId = m.SenderId
	}))

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed on the 'notEqual' tag")

	mockClient.AssertNotCalled(t, "Search", mock.Anything, mock.Anything)
}

func TestMessageValidator(t *testing.T) {
	tests := []struct {
		name              string
		message           *dto.MessageCreate
		mockSearchResp    *employee.SearchResponse
		mockSearchErr     error
		expectErrContains string
	}{
		{
			name:           "Valid message",
			message:        NewMessageCreate(),
			mockSearchResp: &employee.SearchResponse{}, // успешный ответ
		},
		{
			name:              "Search employee returns error",
			message:           NewMessageCreate(),
			mockSearchErr:     errors.New("employee check failed"),
			expectErrContains: "employee check failed",
		},
		{
			name:              "Search returns nil response",
			message:           NewMessageCreate(),
			mockSearchResp:    nil,
			expectErrContains: "not found employees",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockEmployeeGrpcClient := mocks.NewEmployeeGrpcClient(t)

			if tc.mockSearchErr != nil {
				mockEmployeeGrpcClient.
					On("Search", mock.Anything, mock.Anything).
					Return(nil, tc.mockSearchErr).Once()
			} else {
				mockEmployeeGrpcClient.
					On("Search", mock.Anything, mock.Anything).
					Return(tc.mockSearchResp, nil).Once()
			}

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
