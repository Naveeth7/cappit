package util

import "github.com/google/uuid"

type ErrorResponse struct {
	Message    string
	StatusCode int
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}
}

func ValidateUUID(id string) error {
	err := uuid.Validate(id)
	if err != nil {
		return err
	}
	return err
}
