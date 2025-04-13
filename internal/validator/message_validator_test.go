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

func TestMessageValidator_InputValidation(t *testing.T) {
	t.Run("Sender equals receiver", func(t *testing.T) {
		mockClient := mocks.NewEmployeeGrpcClient(t)
		validator := NewMessageValidator(clientValidator.NewEmployeeValidator(mockClient))

		msg := newTestMessage(func(m *dto.MessageCreate) {
			m.ReceiverId = m.SenderId
		})

		err := validator.Validate(msg)

		require.Error(t, err)
		assert.EqualError(t, err, "sender and receiver must be different")
		mockClient.AssertNotCalled(t, "Search")
	})

	t.Run("Empty message text", func(t *testing.T) {
		mockClient := mocks.NewEmployeeGrpcClient(t)
		validator := NewMessageValidator(clientValidator.NewEmployeeValidator(mockClient))

		msg := newTestMessage(func(m *dto.MessageCreate) {
			m.Text = ""
		})

		err := validator.Validate(msg)

		require.Error(t, err)
		assert.EqualError(t, err, "text is required")
		mockClient.AssertNotCalled(t, "Search")
	})
}

func TestMessageValidator_EmployeeValidation(t *testing.T) {
	testCases := []struct {
		name          string
		mockResponse  *employee.SearchResponse
		mockError     error
		expectedError string
	}{
		{
			name: "Valid employees",
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
			expectedError: "employee validation failed: employee check failed: connection error",
		},
		{
			name: "Employee not found",
			mockResponse: &employee.SearchResponse{
				Employees: []*employee.Employee{{Id: testSenderID}},
			},
			expectedError: "employee validation failed: one or more employees do not exist",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := mocks.NewEmployeeGrpcClient(t)
			expectedRequest := &employee.SearchRequest{
				Ids: []string{testSenderID, testReceiverID},
			}

			mockClient.On("Search", mock.Anything, expectedRequest).
				Return(tc.mockResponse, tc.mockError).
				Once()

			validator := NewMessageValidator(
				clientValidator.NewEmployeeValidator(mockClient),
			)

			err := validator.Validate(newTestMessage())

			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)
		})
	}
}
