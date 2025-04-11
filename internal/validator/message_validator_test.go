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

const (
	testChatID     = "00000000-0000-0000-0000-000000000001"
	testSenderID   = "00000000-0000-0000-0000-000000000002"
	testReceiverID = "00000000-0000-0000-0000-000000000003"
	testMessageTxt = "Test message"
)

func newTestMessage(modifiers ...func(*dto.MessageCreate)) *dto.MessageCreate {
	msg := &dto.MessageCreate{
		ChatId:     uuid.MustParse(testChatID),
		SenderId:   uuid.MustParse(testSenderID),
		ReceiverId: uuid.MustParse(testReceiverID),
		Text:       testMessageTxt,
	}

	for _, modify := range modifiers {
		modify(msg)
	}
	return msg
}

func newTestSearchRequest() *employee.SearchRequest {
	return &employee.SearchRequest{
		Ids: []string{testSenderID, testReceiverID},
	}
}

func TestMessageValidator_SenderEqualsReceiver(t *testing.T) {
	mockClient := mocks.NewEmployeeGrpcClient(t)
	validator := NewMessageValidator(clientValidator.NewEmployeeValidator(mockClient))

	msg := newTestMessage(func(m *dto.MessageCreate) {
		m.ReceiverId = m.SenderId
	})

	err := validator.Validate(msg)

	require.Error(t, err)
	assert.EqualError(t, err, "sender and receiver must be different")
	mockClient.AssertNotCalled(t, "Search")
}

func TestMessageValidator(t *testing.T) {
	type validatorTestCase struct {
		name          string
		mockResponse  *employee.SearchResponse
		mockError     error
		expectedError error
	}

	testCases := []validatorTestCase{
		{
			name: "Valid message",
			mockResponse: &employee.SearchResponse{
				Employees: []*employee.Employee{
					{Id: testSenderID},
					{Id: testReceiverID},
				},
			},
		},
		{
			name:          "gRPC connection error",
			mockError:     errors.New("connection error"),
			expectedError: errors.New("employee validation failed: employee check failed: connection error"),
		},
		{
			name: "Employee not found",
			mockResponse: &employee.SearchResponse{
				Employees: []*employee.Employee{{Id: testSenderID}},
			},
			expectedError: errors.New("employee validation failed: one or more employees do not exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := mocks.NewEmployeeGrpcClient(t)
			mockClient.On("Search", mock.Anything, newTestSearchRequest()).
				Return(tc.mockResponse, tc.mockError).
				Once()

			validator := NewMessageValidator(
				clientValidator.NewEmployeeValidator(mockClient),
			)

			err := validator.Validate(newTestMessage())

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
